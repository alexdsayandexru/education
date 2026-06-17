package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDbUrl(host, port, name, user, password, sslmode string) string {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, name, sslmode)
	return connStr
}

func GetPool(host, port, name, user, password, sslmode string) (*pgxpool.Pool, error) {
	return GetPool2(GetDbUrl(host, port, name, user, password, sslmode))
}

func GetPool2(dbUrl string) (*pgxpool.Pool, error) {
	if db, err := pgxpool.New(context.Background(), dbUrl); err != nil {
		return nil, err
	} else {
		return db, db.Ping(context.Background())
	}
}
