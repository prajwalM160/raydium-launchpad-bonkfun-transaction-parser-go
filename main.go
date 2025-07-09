package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type tokenAuth struct {
	token string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{"x-token": t.token}, nil
}

func (tokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("❌ PRIVATE_KEY environment variable not set")
	}

	grpcEndpoint := os.Getenv("GRPC_ENDPOINT")
	if grpcEndpoint == "" {
		log.Fatal("❌ GRPC_ENDPOINT environment variable not set")
	}

	grpcAuthToken := os.Getenv("GRPC_AUTH_TOKEN")
	if grpcAuthToken == "" {
		log.Fatal("❌ GRPC_AUTH_TOKEN environment variable not set")
	}

}
