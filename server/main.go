package main

import (
	"context"
	"fmt"
	"log"
	"net"

	go_example "github.com/go-client-server/protobuf-src"

	"google.golang.org/grpc"
)

const (
	listenHost = "localhost"
	listenPort = "50001"
)

type grpcServer struct {
}

//
// Pay Attention to the first letter in the function name being capitalized in the *.pb.go file
//
func (s *grpcServer) ReverseIt(ctx context.Context, in *go_example.StringRequest) (*go_example.StringResponse, error) {
	myString := in.GetContent()
	log.Printf("Received: %v", myString)
	r := []rune(myString)
	rLength := len(r)
	for i, j := 0, rLength-1; i < rLength/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	log.Printf("Massaged: %v", string(r))
	return &go_example.StringResponse{Content: string(r)}, nil
}

func main() {
	fmt.Println("We are running")

	createServer()
}

func createServer() {

	lis, err := net.Listen("tcp", listenHost+":"+listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srvr := grpc.NewServer()

	go_example.RegisterStringManipServer(srvr, &grpcServer{})

	err = srvr.Serve(lis)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
