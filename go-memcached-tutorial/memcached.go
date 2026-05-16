package main

import (
	"bytes"
	"encoding/gob"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Client struct {
	client *memcache.Client
}

func NewMemcached() (*Client, error) {
	client := memcache.New(os.Getenv("MEMCACHED"))
	if err := client.Ping(); err != nil {
		logger.Error("memcached ping failed", "error", err)
		return nil, err
	}

	client.Timeout = 100 * time.Millisecond
	client.MaxIdleConns = 100

	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetName(nconst string) (Name, error) {
	item, err := c.client.Get(nconst)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			logger.Info("memcached lookup missed", "nconst", nconst)
		} else {
			logger.Warn("memcached lookup failed", "nconst", nconst, "error", err)
		}
		return Name{}, err
	}

	b := bytes.NewReader(item.Value)

	var res Name

	if err := gob.NewDecoder(b).Decode(&res); err != nil {
		logger.Warn("memcached decode failed", "nconst", nconst, "error", err)
		return Name{}, err
	}

	logger.Info("memcached lookup hit", "nconst", nconst)
	return res, nil
}

func (c *Client) SetName(n Name) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		logger.Warn("memcached encode failed", "nconst", n.NConst, "error", err)
		return err
	}

	if err := c.client.Set(&memcache.Item{
		Key:        n.NConst,
		Value:      b.Bytes(),
		Expiration: int32(time.Now().Add(25 * time.Second).Unix()),
	}); err != nil {
		logger.Warn("memcached set failed", "nconst", n.NConst, "error", err)
		return err
	}

	logger.Info("memcached item stored", "nconst", n.NConst)
	return nil
}
