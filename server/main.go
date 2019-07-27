package main

import (
	"fmt"
	pb "github.com/rai-prashanna/gateway/template"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/big"
	"net"
	"strconv"
)
/* Variable Declaration */
var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers.
// Range: 0 through 18446744073709551615.
var i int = 1
var n int

const (
	port = ":10000"
)

type server struct{}

func (s *server) SendGet(ctx context.Context, in *pb.TemplateRequest) (*pb.TemplateResponse, error) {
	var fact big.Int
	return &pb.TemplateResponse{Message: "Received GET method " + fact.MulRange(1,in.Num).String()}, nil
}

func (s *server) SendPost(ctx context.Context, in *pb.TemplateRequest) (*pb.TemplateResponse, error) {
	return &pb.TemplateResponse{Message: "Received POST method " + in.Name}, nil
}
/*     function declaration        */
func factorial(n int) uint64 {
	if(n < 0){
		fmt.Print("Factorial of negative number doesn't exist.")
	}else{
		for i:=1; i<=n; i++ {
			factVal *= uint64(i)  // mismatched types int64 and int
		}

	}
	return factVal  /* return from function*/
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
