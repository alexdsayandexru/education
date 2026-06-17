package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	flipt "github.com/open-feature/go-sdk-contrib/providers/flipt/pkg/provider"
	"github.com/open-feature/go-sdk/openfeature"
)

func client1() {
	err := openfeature.SetProviderAndWait(flipt.NewProvider())
	if err != nil {
		panic(err)
	}

	client := openfeature.NewClient("my-app")

	for i := 0; i < 10; i++ {
		if value, err := client.StringValue(context.Background(), "colors", "white", openfeature.NewEvaluationContext(
			uuid.New().String(),
			map[string]any{
				"role": "admin",
			},
		)); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(value)
		}
	}
}
