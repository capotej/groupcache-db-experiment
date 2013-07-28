package main

import (
	"github.com/capotej/groupcachedb/server"
	"github.com/capotej/groupcachedb/slowdb"
	"github.com/golang/groupcache"
)

func main() {

	db := slowdb.NewSlowDB()

	var stringcache = groupcache.NewGroup("SlowDBCache", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			result := db.Get(key)
			dest.SetBytes([]byte(result))
			return nil
		}))

	server := server.NewServer(stringcache, db)

	server.Start()
}
