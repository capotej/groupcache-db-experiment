package main

import (
	"flag"
	"fmt"
	"github.com/capotej/groupcache-db-experiment/api"
	"github.com/capotej/groupcache-db-experiment/client"
	"net/rpc"
)

func main() {

	var port = flag.String("port", "9001", "frontend port")
	var set = flag.Bool("set", false, "doing a set?")
	var get = flag.Bool("get", false, "doing a get?")
	var cget = flag.Bool("cget", false, "doing a get?")
	var key = flag.String("key", "foo", "key to get")
	var value = flag.String("value", "bar", "value to set")
	flag.Parse()

	client := new(client.Client)
	if *cget {
		client, err := rpc.DialHTTP("tcp", "localhost:"+*port)
		if err != nil {
			fmt.Printf("error %s", err)
		}
		args := &api.Load{*key}
		var reply api.ValueResult
		err = client.Call("Frontend.Get", args, &reply)
		if err != nil {
			fmt.Printf("error %s", err)
		}
		fmt.Println(string(reply.Value))
		return
	}

	if *get {
		var reply = client.Get(*key)
		fmt.Println(reply)
		return
	}

	if *set {
		client.Set(*key, *value)
		return
	}

	flag.PrintDefaults()
	return

}
