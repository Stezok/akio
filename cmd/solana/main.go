package main

import (
	"context"
	"log"
	"os"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	balance, err := c.GetBalance(context.Background(), "Gndmrf2q1KwzcbTZPF1iAzxf49wXW4GDLDaJvZELLEg3")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Your balance: %d SOL", balance/1e9)

	os.Exit(0)

	account, err := types.AccountFromBase58("5HR9qjK186U1Zv4kCxhV4TonjJQj9K2KXKqvjJiX377Bj7ghzqCyS5e1EFJzfcmuw1CjsLkjhAxc4PTvAbvLQbWs")
	if err != nil {
		log.Fatal(err)
	}

	txhash, err := c.RequestAirdrop(
		context.Background(),
		account.PublicKey.ToBase58(),
		1e9, // 1 SOL = 10^9 lamports
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("TxHash: %s", txhash)
}
