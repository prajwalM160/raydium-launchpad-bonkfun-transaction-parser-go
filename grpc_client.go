package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type GeyserClientWrapper struct {
	Client pb.GeyserClient
	Conn   *grpc.ClientConn
}

func NewGeyserClient(grpcEndpoint, grpcAuthToken string) (*GeyserClientWrapper, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(tokenAuth{token: grpcAuthToken}),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024*1024*1024),
			grpc.MaxCallSendMsgSize(1024*1024*1024),
		),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * 1e9, // 10 seconds
			Timeout:             5 * 1e9,  // 5 seconds
			PermitWithoutStream: true,
		}),
	}

	log.Printf("🔌 Connecting to Geyser: %s", grpcEndpoint)
	conn, err := grpc.Dial(grpcEndpoint, opts...)
	if err != nil {
		return nil, err
	}
	log.Println("✅ gRPC Connection Established.")
	client := pb.NewGeyserClient(conn)
	return &GeyserClientWrapper{Client: client, Conn: conn}, nil
}

func (g *GeyserClientWrapper) SubscribePumpFun(ctx context.Context, programID string) (pb.Geyser_SubscribeClient, error) {
	voteFilter := false
	failedFilter := false

	subReq := &pb.SubscribeRequest{
		Transactions: map[string]*pb.SubscribeRequestFilterTransactions{
			"pump_fun_subscription": {
				Vote:           voteFilter,
				Failed:         failedFilter,
				AccountInclude: []string{programID},
			},
		},
		TransactionsStatus: map[string]*pb.SubscribeRequestFilterTransactions{
			"pump_fun_status": {
				Vote:           voteFilter,
				Failed:         failedFilter,
				AccountInclude: []string{programID},
			},
		},
		Commitment: pb.CommitmentLevel_PROCESSED,
	}

	stream, err := g.Client.Subscribe(ctx)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(subReq); err != nil {
		return nil, err
	}
	log.Println("✅ Subscribed. Waiting for 'create' transactions...")
	return stream, nil
}
