package main

import "fmt"
import "github.com/capotej/groupcachedb/server"
import "net/rpc"
import "os"

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	args := &server.Store{"foo", "bar"}
	var reply int
	err = client.Call("Server.Set", args, &reply)
	if err != nil {
		fmt.Printf("error %s", err)
	}

}
