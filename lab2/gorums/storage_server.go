package gorums

import (
	"log"
	"net"
	"sync"

	pb "dat520/info/lab2/gorums/proto"

	"github.com/relab/gorums"
	"google.golang.org/protobuf/types/known/emptypb"
)

// StorageServer implements the StorageService interface defined in storage.proto.
type StorageServer struct {
	sync.Mutex
	data []string
	// TODO: Add fields if necessary
}

// NewStorageServer creates a new StorageServer.
func NewStorageServer() *StorageServer {
	return &StorageServer{
		data: make([]string, 0),
	}
}

// Start starts the server and returns the address the server is listening on.
// The function should be non-blocking.
func (s *StorageServer) Start() (string, error) {
	lis, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		return "", err
	}
	gorumsSrv := gorums.NewServer()
	pb.RegisterStorageServiceServer(gorumsSrv, s)
	go func() {
		if err := gorumsSrv.Serve(lis); err != nil {
			log.Printf("failed to start serving: %v", err)
		}
	}()
	return lis.Addr().String(), nil
}

// Returns the data slice on this server.
func (s *StorageServer) GetData() []string {
	s.Lock()
	defer s.Unlock()
	return s.data
}

// Write appends the value to the data slice on this server.
func (s *StorageServer) Write(ctx gorums.ServerCtx, request *pb.WriteRequest) (*emptypb.Empty, error) {
	s.Lock()
	defer s.Unlock()
	s.data = append(s.data, request.Value)
	return &emptypb.Empty{}, nil
}

// Read returns the data slice on this server.
func (s *StorageServer) Read(ctx gorums.ServerCtx, request *emptypb.Empty) (*pb.ReadResponse, error) {
	s.Lock()
	defer s.Unlock()
	return &pb.ReadResponse{Values: s.data}, nil
}
