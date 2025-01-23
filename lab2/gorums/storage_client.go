package gorums

import (
	"context"
	"fmt"
	"math/rand/v2"
	"slices"
	"time"

	pb "dat520/info/lab2/gorums/proto"

	"github.com/relab/gorums"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

// StorageClient is a client for the StorageService defined in storage.proto.
type StorageClient struct {
	conf *pb.Configuration
}

// NewStorageClient creates a new StorageClient with the provided srvAddrs as the configuration.
func NewStorageClient(srvAddrs []string) (*StorageClient, error) {
	if len(srvAddrs) == 0 {
		return nil, fmt.Errorf("no server addresses provided")
	}
	mgr := pb.NewManager(gorums.WithGrpcDialOptions(
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	))
	cfg, err := mgr.NewConfiguration(&qspec{len(srvAddrs)}, gorums.WithNodeList(srvAddrs))
	if err != nil {
		return nil, fmt.Errorf("unable to create configuration: %v", err)
	}
	return &StorageClient{
		conf: cfg,
	}, nil
}

// WriteValue writes the provided value to a random server.
func (sc *StorageClient) WriteValue(value string) error {
	allNodes := sc.conf.Nodes()
	node := allNodes[rand.IntN(len(allNodes))]
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	_, err := node.Write(ctx, &pb.WriteRequest{Value: value})
	if err != nil {
		return fmt.Errorf("unable to write value: %v", err)
	}
	return nil
}

// ReadValues reads all values from all servers and returns them in sorted order.
func (sc *StorageClient) ReadValues() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	resp, err := sc.conf.Read(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("unable to read values: %w", err)
	}
	return resp.Values, nil
}

type qspec struct {
	numServers int
}

func (qs *qspec) ReadQF(in *emptypb.Empty, replies map[uint32]*pb.ReadResponse) (*pb.ReadResponse, bool) {
	if len(replies) != qs.numServers {
		return nil, false
	}
	values := make([]string, 0, len(replies))
	for _, r := range replies {
		values = append(values, r.Values...)
	}
	return &pb.ReadResponse{Values: slices.Sorted(slices.Values(values))}, true
}
