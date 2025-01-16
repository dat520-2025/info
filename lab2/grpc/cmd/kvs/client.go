package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	pb "dat520/info/lab2/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func client(n int, endpoint string) {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to server:", endpoint)

	client := pb.NewKeyValueServiceClient(conn)

	for i := range n {
		req := &pb.InsertRequest{Key: strconv.Itoa(i), Value: fmt.Sprintf("value%d", i)}
		resp, err := client.Insert(context.Background(), req)
		if err != nil {
			log.Fatalf("Failed to insert key %d: %v", i, err)
		}
		log.Printf("Inserted key %d: %v\n", i, resp)
	}

	for i := range n {
		resp, err := client.Lookup(context.Background(), &pb.LookupRequest{Key: strconv.Itoa(i)})
		if err != nil {
			log.Fatalf("Failed to lookup key %d: %v", i, err)
		}
		log.Printf("Looked up key %d: %v\n", i, resp)
	}

	resp, err := client.Keys(context.Background(), &pb.KeysRequest{})
	if err != nil {
		log.Fatalf("Failed to get all keys: %v", err)
	}
	log.Printf("All keys: %v\n", resp)
}
