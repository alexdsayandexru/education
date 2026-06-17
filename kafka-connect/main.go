package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

func parseRecord(record map[string]interface{}) (fields []string, values []interface{}) {
	payload := record["payload"].(map[string]interface{})

	for _, field_ := range record["schema"].(map[string]interface{})["fields"].([]interface{}) {
		field := field_.(map[string]interface{})
		name := field["field"].(string)
		fields = append(fields, name)
		if payload[name] != nil && field["name"] != nil {
			if field["name"].(string) == "io.debezium.time.MicroTimestamp" {
				values = append(values, time.UnixMicro(int64(payload[name].(float64))).UTC())
			} else if field["name"].(string) == "io.debezium.time.Date" {
				values = append(values, time.Unix(int64(payload[name].(float64))*86400, 0).UTC())
			} else {
				values = append(values, payload[name])
			}
		} else {
			values = append(values, payload[name])
		}
	}
	return
}

func getInsertQuery(table string, names []string) string {
	var fields string
	var values string

	for i, name := range names {
		if fields != "" {
			fields += ","
		}
		fields += name
		if values != "" {
			values += ","
		}
		values += fmt.Sprintf("$%d", i+1)
	}
	return fmt.Sprintf("insert into %s (%s) values (%s)", table, fields, values)
}

func getSelectQuery(table string, idName string) string {
	return fmt.Sprintf("select count(*) from %s where %s = $1", table, idName)
}

func getUpdateQuery(table string, names []string) string {
	var fields string
	for i, name := range names {
		if i == 0 {
			continue
		}
		if fields != "" {
			fields += ","
		}
		fields += name + "=" + fmt.Sprintf("$%d", i+1)
	}
	return fmt.Sprintf("update %s set %s where %s = $1", table, fields, names[0])
}

func main() {
	jsonFile, err := os.Open("obj.json")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = jsonFile.Close()
	}()

	byteValue, _ := io.ReadAll(jsonFile)
	var record interface{}

	if err = json.Unmarshal(byteValue, &record); err != nil {
		fmt.Println(err)
		return
	} else {
		if pool, err := GetPool("localhost", "5434", "master1", "postgres", "123", "disable"); err != nil {
			fmt.Println(err)
			return
		} else {
			fields, values := parseRecord(record.(map[string]interface{}))

			query := getSelectQuery("users", fields[0])
			row := pool.QueryRow(context.Background(), query, values[0])
			var count int
			if err := row.Scan(&count); err != nil {
				fmt.Println(err)
				return
			} else if count == 0 {
				query = getInsertQuery("users", fields)
				if _, err := pool.Exec(context.Background(), query, values...); err != nil {
					fmt.Println(err)
					return
				}
			} else if count == 1 {
				query = getUpdateQuery("users", fields)
				if _, err := pool.Exec(context.Background(), query, values...); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
