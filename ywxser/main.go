package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	ywx "grpcstudy/idl"
	"grpcstudy/utils"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	ywx.UnimplementedTeserServer
}

func (s *server) GetPid(ctx context.Context, in *ywx.PidRequest) (*ywx.PidReply, error) {
	id := &ywx.PidReply{OutId: in.GetInId()}
	return id, nil
}
func (s *server) GetIP(ctx context.Context, in *ywx.IPRequest) (*ywx.IPReply, error) {
	ip := utils.GetLocalIp()
	res := &ywx.IPReply{IP: ip}
	return res, nil
}
func main() {
	fmt.Println("-------grpc server start--------\n")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port)) //监听端口
	if err != nil {
		log.Fatalf("failed to listen: %v", listen)
	}
	s := grpc.NewServer() //create 一个没有注册的rpc服务器
	ywx.RegisterTeserServer(s, &server{})
	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
