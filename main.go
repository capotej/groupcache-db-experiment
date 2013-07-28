package main

import (
	"github.com/capotej/groupcachedb/server"
	"github.com/capotej/groupcachedb/slowdb"
	"github.com/golang/groupcache"
	"net/http"
)

func main() {

	db := slowdb.NewSlowDB()

	peers := groupcache.NewHTTPPool("http://localhost:8001")

	var stringcache = groupcache.NewGroup("SlowDBCache", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			result := db.Get(key)
			dest.SetBytes([]byte(result))
			return nil
		}))

	go http.ListenAndServe("127.0.0.1:8001", http.HandlerFunc(peers.ServeHTTP))

	server := server.NewServer(stringcache, db)

	server.Start(":8080")

}
