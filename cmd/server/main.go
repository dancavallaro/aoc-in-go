package main

import (
	"aoc-in-go/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type aocSolverService struct {
	proto.UnimplementedAOCSolverServer
}

func (svc aocSolverService) Solve(_ context.Context, req *proto.SolveRequest) (*proto.SolveResponse, error) {
	fmt.Printf("Received Solve request: %s\n", req)
	return &proto.SolveResponse{Result: "lol"}, nil
}

func main() {
	port := flag.Int("port", 42069, "TCP port to listen on")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	proto.RegisterAOCSolverServer(server, aocSolverService{})
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
