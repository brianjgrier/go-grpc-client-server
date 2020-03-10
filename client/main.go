package main

import (
	"context"
	"fmt"
	"log"
	"time"

	go_example "github.com/go-client-server/protobuf-src"
	includes "github.com/go-client-server/imports"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(includes.ListenHost+":"+includes.ListenPort, grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))

	if err != nil {
		log.Fatal("Could not connect to server")
		return
	}
	defer conn.Close()

	client := go_example.NewStringManipClient(conn)

	response, err := client.ReverseIt(context.Background(), &go_example.StringRequest{
		Content: "Mares Eat Oats",
	})

	if err != nil {
		log.Fatal("Could not get response from server")
	}

	fmt.Println(response)
}
