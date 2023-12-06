package main

import (
	"aoc-in-go/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = flag.String("address", "localhost:42069", "address of server")
	year    = flag.Int("year", 0, "year of puzzle")
	day     = flag.Int("day", 0, "day of puzzle")
	part2   = flag.Bool("part2", false, "whether to solve part 2 (otherwise, part 1)")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewAOCSolverClient(conn)
	response, err := client.Solve(context.TODO(), &proto.SolveRequest{
		Year: int32(*year), Day: int32(*day), Part2: *part2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %s\n", response.Result)
}
