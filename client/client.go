package main

import "fmt"

import "github.com/capotej/groupcachedb/server"
import "net/rpc"
import "os"

func main() {

	switch os.Args[1] {
	case "get":
		var key = os.Args[2]
		client, err := rpc.DialHTTP("tcp", "localhost:8080")
		if err != nil {
			fmt.Printf("error %s", err)
		}
		args := &server.Load{key}
		var reply server.ValueResult
		err = client.Call("Server.Get", args, &reply)
		if err != nil {
			fmt.Printf("error %s", err)
		}
		fmt.Println(reply)
		return

	case "set":
		var key = os.Args[2]
		var value = os.Args[3]

		client, err := rpc.DialHTTP("tcp", "localhost:8080")
		if err != nil {
			fmt.Printf("error %s", err)
		}
		args := &server.Store{key, value}
		var reply int
		err = client.Call("Server.Set", args, &reply)
		if err != nil {
			fmt.Printf("error %s", err)
		}

		return
	}

	fmt.Println("please use set or get")
	return

}
