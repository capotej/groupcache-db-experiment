package slowdb

// A purposefully slow key/value database

import (
	"fmt"
	"time"
)

type SlowDB struct {
	data map[string]string
}

func (db *SlowDB) Get(key string) string {
	time.Sleep(time.Duration(300) * time.Millisecond)
	fmt.Printf("getting %s\n", key)
	return db.data[key]
}

func (db *SlowDB) Set(key string, value string) {
	fmt.Printf("setting %s to %s\n", key, value)
	db.data[key] = value
}

func NewSlowDB() *SlowDB {
	ndb := new(SlowDB)
	ndb.data = make(map[string]string)
	return ndb
}
