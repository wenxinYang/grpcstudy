package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	ywx "grpcstudy/idl"
	"log"
	"time"
)

const (
	defaultName = "ywx-test"
)

var (
	//Input_flagvar string
	inId int
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	//name = flag.String("name", defaultName, "Name to greet")
)

func init() {
	flag.IntVar(&inId, "inId", 123, "input id")
}
func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := ywx.NewTeserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pid, err := client.GetPid(ctx, &ywx.PidRequest{InId: int32(inId)})
	if err != nil {
		log.Fatalf("get pid fail")
	}
	fmt.Printf("pid = %v\n", pid.GetOutId())

	ip, _ := client.GetIP(ctx, &ywx.IPRequest{})
	fmt.Println("ip = ", ip.GetIP())
}
