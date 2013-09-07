package client

// Client for dbserver/slowdb

import (
	"fmt"
	"github.com/capotej/groupcache-db-experiment/api"
	"net/rpc"
)

type Client struct{}

func (c *Client) Get(key string) string {
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	args := &api.Load{key}
	var reply api.ValueResult
	err = client.Call("Server.Get", args, &reply)
	if err != nil {
		fmt.Printf("error %s", err)
	}
	return string(reply.Value)
}

func (c *Client) Set(key string, value string) {
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	args := &api.Store{key, value}
	var reply int
	err = client.Call("Server.Set", args, &reply)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}
