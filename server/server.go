package server

import (
	"fmt"
	"github.com/capotej/groupcachedb/slowdb"
	"github.com/golang/groupcache"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
	cacheGroup *groupcache.Group
	db         *slowdb.SlowDB
}

//a "Load" rpc instruction
type Load struct {
	Key string
}

//a "Store" rpc instruction
type Store struct {
	Key   string
	Value string
}

type NullResult int

type ValueResult struct {
	Value string
}

func (s *Server) Get(args *Load, reply *ValueResult) error {
	var data []byte
	err := s.cacheGroup.Get(nil, args.Key,
		groupcache.AllocatingByteSliceSink(&data))

	reply.Value = string(data)
	return err
}

func (s *Server) Set(args *Store, reply *NullResult) error {
	s.db.Set(args.Key, args.Value)
	*reply = 0
	return nil
}

func NewServer(cacheGroup *groupcache.Group, db *slowdb.SlowDB) *Server {
	server := new(Server)
	server.db = db
	server.cacheGroup = cacheGroup
	return server
}

func (s *Server) Start() {

	rpc.Register(s)

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		fmt.Println("fatal")
	}

	http.Serve(l, nil)
}
