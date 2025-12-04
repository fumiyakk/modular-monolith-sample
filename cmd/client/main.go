package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fumiyakk/modular-monolith-sample/internal/client"
)

func main() {
	c := client.NewClient("http://localhost:8080")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ユーザーと契約を作成
	result, err := c.CreateUser(ctx, "Test User")
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	fmt.Printf("Created user (ID: %s) with contract (ID: %s)\n", result.UserID, result.ContractID)

	// ユーザー情報を取得
	user, err := c.GetUser(ctx, result.UserID)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	fmt.Printf("Got user: ID=%s, Name=%s\n", user.Id, user.Name)

	// 契約情報を取得
	contract, err := c.GetContract(ctx, result.ContractID)
	if err != nil {
		log.Fatalf("Failed to get contract: %v", err)
	}
	fmt.Printf("Got contract: ID=%s, UserID=%s, Status=%v\n",
		contract.Id,
		contract.UserId,
		contract.Status,
	)
}
