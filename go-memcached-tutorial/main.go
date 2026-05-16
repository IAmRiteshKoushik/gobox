package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/mux"
)

type Name struct {
	NConst    string `json:"nconst"`
	Name      string `json:"name"`
	BirthYear string `json:"birthYear"`
	DeathYear string `json:"deathYear"`
}

type Error struct {
	Message string `json:"error"`
}

func main() {
	db, err := newDB()
	if err != nil {
		logger.Error("failed to initialize database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	mc, err := NewMemcached()
	if err != nil {
		logger.Error("failed to initialize memcached client", "error", err)
		os.Exit(1)
	}

	router := mux.NewRouter()

	renderJSON := func(w http.ResponseWriter, val interface{}, statusCode int) {
		w.WriteHeader(statusCode)
		_ = json.NewEncoder(w).Encode(val)
	}

	router.HandleFunc("/names/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		reqLogger := logger.With(
			"method", r.Method,
			"path", r.URL.Path,
			"nconst", id,
		)

		reqLogger.Info("handling name lookup")

		val, err := mc.GetName(id)
		if err == nil { // Cache hit
			reqLogger.Info("serving response from cache", "status", http.StatusOK)
			renderJSON(w, &val, http.StatusOK)
			return
		}
		if errors.Is(err, memcache.ErrCacheMiss) {
			reqLogger.Info("cache miss")
		} else {
			reqLogger.Warn("cache read failed", "error", err)
		}

		// DB hit
		name, err := db.FindByNConst(id)
		if err != nil { // miss
			reqLogger.Error("database lookup failed", "error", err, "status", http.StatusInternalServerError)
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		if err := mc.SetName(name); err != nil {
			reqLogger.Warn("failed to repopulate cache", "error", err)
		}

		reqLogger.Info("serving response from database", "status", http.StatusOK)
		renderJSON(w, &name, http.StatusOK)
	})

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("starting server", "addr", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server stopped unexpectedly", "error", err)
		os.Exit(1)
	}
}
