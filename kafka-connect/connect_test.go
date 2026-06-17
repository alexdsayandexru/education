package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"sync"
	"testing"
	"time"
)

func randText(tag string) string {
	return tag + "-" + uuid.New().String()
}

func FillTable(dbName string, countRecords int) error {
	db, err := pgx.Connect(context.Background(), GetDbUrl("localhost", "5434", dbName, "postgres", "123", "disable"))
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close(context.Background())
	}()

	batch := pgx.Batch{}

	for i := 0; i < countRecords; i++ {
		query := "INSERT INTO public.users(gid, first_name, last_name, patronymic, birth_date, nickname, gender, created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
		batch.Queue(query, uuid.New().String(), randText(dbName), randText(dbName), randText(dbName), time.Now(), dbName, "m", time.Now())
	}

	br := db.SendBatch(context.Background(), &batch)
	defer func() {
		_ = br.Close()
	}()
	_, err = br.Exec()
	return err
}

func TestSynchronization(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		name := "idp"
		if i%2 == 0 {
			name += "2"
		}
		go func(name string) {
			if err := FillTable(name, 10); err != nil {
				t.Error(err)
			}
			wg.Done()
		}(name)
	}
	wg.Wait()
}
