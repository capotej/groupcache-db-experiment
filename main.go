package main

import (
	"flag"
	"fmt"
	"github.com/capotej/groupcachedb/server"
	"github.com/capotej/groupcachedb/slowdb"
	"github.com/golang/groupcache"
	"net/http"
)

func main() {

	var port = flag.String("port", "8001", "groupcache port")
	var master = flag.Bool("master", false, "are we rpc master")
	var rpcport = flag.String("rpc", "8080", "port for rpc master")
	flag.Parse()

	db := slowdb.NewSlowDB()

	peers := groupcache.NewHTTPPool("http://localhost:" + *port)

	var stringcache = groupcache.NewGroup("SlowDBCache", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			result := db.Get(key)
			dest.SetBytes([]byte(result))
			return nil
		}))

	fmt.Println(stringcache)

	peers.Set("http://localhost:8001", "http://localhost:8002", "http://localhost:8003")

	if *master {
		fmt.Println("master starting on " + *rpcport)
		go http.ListenAndServe("127.0.0.1:"+*port, http.HandlerFunc(peers.ServeHTTP))
		server := server.NewServer(stringcache, db)
		server.Start(":" + *rpcport)
	} else {
		fmt.Println("slave starting on " + *port)
		http.ListenAndServe("127.0.0.1:"+*port, http.HandlerFunc(peers.ServeHTTP))
	}

}
