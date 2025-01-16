package kvstore

import (
	"context"
	"log"
	"maps"
	"slices"
	"sync"

	pb "dat520/info/lab2/grpc/proto"
)

type keyValueServicesServer struct {
	mu sync.Mutex
	kv map[string]string
	// this must be included in implementers of the pb.KeyValueServicesServer interface
	pb.UnimplementedKeyValueServiceServer
}

// NewKeyValueServicesServer returns an initialized KeyValueServicesServer
func NewKeyValueServicesServer() *keyValueServicesServer {
	return &keyValueServicesServer{
		kv: make(map[string]string),
	}
}

// Insert inserts a key-value pair from the request into the server's map, and
// returns a response to the client indicating whether or not the insert was successful.
func (s *keyValueServicesServer) Insert(ctx context.Context, req *pb.InsertRequest) (*pb.InsertResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.kv[req.Key] = req.Value
	log.Printf("Inserted key-value pair: %v\n", req)
	return &pb.InsertResponse{Success: true}, nil
}

// Lookup returns a response to containing the value corresponding to the request's key.
// If the key is not found, the response's value is empty.
func (s *keyValueServicesServer) Lookup(ctx context.Context, req *pb.LookupRequest) (*pb.LookupResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val := s.kv[req.Key]
	return &pb.LookupResponse{Value: val}, nil
}

// Keys returns a response to containing a slice of all the keys in the server's map.
// The returned slice is sorted.
func (s *keyValueServicesServer) Keys(ctx context.Context, req *pb.KeysRequest) (*pb.KeysResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return &pb.KeysResponse{Keys: slices.Sorted(maps.Keys(s.kv))}, nil
}
