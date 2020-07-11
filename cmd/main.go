package main

import (
	"context"
	"fmt"
	"log"

	"github.com/qushot/singleton-datastore-client/domain/model"
	"github.com/qushot/singleton-datastore-client/registry"
)

func main() {
	reg, err := registry.Get()
	if err != nil {
		log.Fatalf("registry get error: %v", err)
		return
	}

	ctx := context.Background()
	user := &model.User{
		Name: "Qushot",
		Age:  26,
	}

	result, err := reg.GetUserRepository().Create(ctx, user)
	if err != nil {
		log.Fatalf("user create error: %v", err)
		return
	}

	fmt.Printf("User: %v\n", result)
}
