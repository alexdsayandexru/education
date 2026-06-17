package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/idp"); err != nil {
		fmt.Println("Connect:", err)
	} else {
		_, err := conn.Exec(context.Background(), "listen users")
		if err != nil {
			fmt.Println("listening to channel error:", err)
		}

		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

		notifyChan := make(chan *pgconn.Notification)

		go func() {
			for {
				notification, err := conn.WaitForNotification(context.Background())
				if err != nil {
					fmt.Println("WaitForNotification:", err)
				}
				notifyChan <- notification
			}
		}()

		for {
			select {
			case notification := <-notifyChan:
				fmt.Println(notification)
			case <-signalChan:
				fmt.Println("break")
				return
			}
		}

	}
	fmt.Println("END")
}
