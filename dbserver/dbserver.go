package main

// This is the rpc server that fronts slowdb

import (
	"fmt"
	"github.com/capotej/groupcache-db-experiment/api"
	"github.com/capotej/groupcache-db-experiment/slowdb"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
	db *slowdb.SlowDB
}

func (s *Server) Get(args *api.Load, reply *api.ValueResult) error {
	data := s.db.Get(args.Key)
	reply.Value = string(data)
	return nil
}

func (s *Server) Set(args *api.Store, reply *api.NullResult) error {
	s.db.Set(args.Key, args.Value)
	*reply = 0
	return nil
}

func NewServer(db *slowdb.SlowDB) *Server {
	server := new(Server)
	server.db = db
	return server
}

func (s *Server) Start(port string) {

	rpc.Register(s)

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", port)
	if e != nil {
		fmt.Println("fatal")
	}

	http.Serve(l, nil)
}

func main() {
	db := slowdb.NewSlowDB()
	server := NewServer(db)
	fmt.Println("dbserver starting on localhost:8080")
	server.Start(":8080")
}
