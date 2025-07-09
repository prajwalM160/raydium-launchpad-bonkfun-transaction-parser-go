package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type SolanaClientWrapper struct {
	Client *rpc.Client
}

func NewSolanaClient() *SolanaClientWrapper {
	rpcEndpoint := os.Getenv("SOLANA_RPC_ENDPOINT")
	if rpcEndpoint == "" {
		apiKey := os.Getenv("HELIUS_API_KEY")
		if apiKey == "" {
			log.Fatal("Missing HELIUS_API_KEY or SOLANA_RPC_ENDPOINT")
		}
		rpcEndpoint = fmt.Sprintf("https://pomaded-lithotomies-xfbhnqagbt-dedicated.helius-rpc.com/?api-key=%s", apiKey)
	}

	log.Println("Connecting to Solana RPC:", rpcEndpoint)
	client := rpc.New(rpcEndpoint)
	return &SolanaClientWrapper{Client: client}
}

func (s *SolanaClientWrapper) GetLatestBlockhash(ctx context.Context, commitment rpc.CommitmentType) (*rpc.GetLatestBlockhashResult, error) {
	return s.Client.GetLatestBlockhash(ctx, commitment)
}

func (s *SolanaClientWrapper) SendTransactionWithOpts(ctx context.Context, tx *solana.Transaction, opts rpc.TransactionOpts) (solana.Signature, error) {
	return s.Client.SendTransactionWithOpts(ctx, tx, opts)
}
