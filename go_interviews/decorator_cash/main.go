package main

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
)

type KVDatabase interface {
	Get(key string) (string, error)
	Keys() ([]string, error)
	MGet(keys []string) ([]*string, error)
}

type Cache struct {
	db   KVDatabase
	mu   sync.RWMutex
	data map[string]string

	dataCreated time.Time
}

const invalidationTTL = time.Minute

func NewCache(db KVDatabase) *Cache {
	c := &Cache{
		db:          db,
		data:        make(map[string]string),
		dataCreated: time.Now(),
	}

	go c.startInvalidator()

	return c
}

func (c *Cache) startInvalidator() {
	tt := time.NewTicker(invalidationTTL)

	for {
		select {
		case <-tt.C:
			c.mu.Lock()
			c.data = make(map[string]string)
			c.mu.Unlock()
		}
	}

}

func (c *Cache) Get(key string) (string, error) {
	c.mu.RLock()
	v, ok := c.data[key]
	c.mu.RUnlock()

	if !ok {
		v, err := c.db.Get(key)
		if err != nil {
			return "", err
		}

		c.mu.Lock()
		c.data[key] = v
		c.mu.Unlock()

	}

	return v, nil

}

func (c *Cache) Keys() ([]string, error) {
	return c.db.Keys()
}

func (c *Cache) MGet(keys []string) ([]*string, error) {
	return c.db.MGet(keys)
}
