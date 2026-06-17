package main

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"time"
)

func main() {
	conn := connect("localhost:9000", "learn_db", "username", "password")
	rows, err := conn.Query("SELECT teacher_id,subject_name,load_date FROM mart_student_lesson limit(10)")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.teacher_id, &u.subject_name, &u.load_date); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(u)
	}
	_ = rows.Close()
}

func connect(host, database, username, password string) *sql.DB {
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{host},
		Auth: clickhouse.Auth{
			Database: database,
			Username: username,
			Password: password,
		},
		/*TLS: &tls.Config{
			InsecureSkipVerify: true,
		},*/
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: time.Second * 30,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "my-app", Version: "0.1"},
			},
		},
	})
	conn.SetMaxIdleConns(5)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(time.Hour)
	return conn
}

type user struct {
	teacher_id   int
	subject_name string
	load_date    time.Time
}

//curl -d '{"metric":{"__name__":"foo","job":"node_exporter"},"values":[0,1,2],"timestamps":[1549891472010,1549891487724,1549891503438]}' -X POST 'http://localhost:8428/api/v1/import'
