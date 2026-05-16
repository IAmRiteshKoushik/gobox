package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func newDB() (*Postgres, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("failed to create postgres connection pool", "error", err)
		return nil, err
	}

	logger.Info("postgres connection pool initialized")
	return &Postgres{
		pool: pool,
	}, nil
}

func (p *Postgres) Close() {
	logger.Info("closing postgres connection pool")
	p.pool.Close()
}

func (p *Postgres) FindByNConst(nconst string) (Name, error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	var res Name

	if err := p.pool.
		QueryRow(context.Background(), query, nconst).
		Scan(&res.NConst, &res.Name, &res.BirthYear, &res.DeathYear); err != nil {
		logger.Warn("postgres lookup failed", "nconst", nconst, "error", err)
		return Name{}, err
	}

	logger.Info("postgres lookup succeeded", "nconst", nconst)
	return res, nil
}
