package main

import (
	"log"
	"os"
)

type Config struct {
	PrivateKey    string
	GrpcEndpoint  string
	GrpcAuthToken string
	RpcEndpoint   string
	HeliusApiKey  string
}

func LoadConfig() Config {
	privateKey := os.Getenv("BUYER_PRIVATE_KEY_PATH")
	if privateKey == "" {
		log.Fatal("❌ BUYER_PRIVATE_KEY_PATH environment variable not set")
	}

	grpcEndpoint := os.Getenv("GRPC_ENDPOINT")
	if grpcEndpoint == "" {
		log.Fatal("❌ GRPC_ENDPOINT environment variable not set")
	}

	grpcAuthToken := os.Getenv("GRPC_AUTH_TOKEN")
	if grpcAuthToken == "" {
		log.Fatal("❌ GRPC_AUTH_TOKEN environment variable not set")
	}

	rpcEndpoint := os.Getenv("SOLANA_RPC_ENDPOINT")
	heliusApiKey := os.Getenv("HELIUS_API_KEY")

	return Config{
		PrivateKey:    privateKey,
		GrpcEndpoint:  grpcEndpoint,
		GrpcAuthToken: grpcAuthToken,
		RpcEndpoint:   rpcEndpoint,
		HeliusApiKey:  heliusApiKey,
	}
}
