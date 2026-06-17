package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.flipt.io/flipt/rpc/flipt/evaluation"
	sdk "go.flipt.io/flipt/sdk/go"
	"go.flipt.io/flipt/sdk/go/http"
)

func client2() {
	transport := http.NewTransport("http://localhost:8181")
	client := sdk.New(transport)

	for i := 0; i < 10; i++ {
		resp, err := client.Evaluation().Boolean(context.Background(), &evaluation.EvaluationRequest{
			NamespaceKey: "default",
			FlagKey:      "factor",
			EntityId:     uuid.New().String(),
			Context: map[string]string{
				"role": "admin",
			},
		})

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp.Enabled)
		}
	}
}
