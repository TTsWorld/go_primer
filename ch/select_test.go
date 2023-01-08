package ch

import (
	"fmt"
	"log"
	"testing"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

type MenuItem struct {
	Id           int64  `db:"id"`
	menu_page_id string `db:"menu_page_id"`
	price        int64  `db:"price"`
	xpos         int64  `db:"xpos"`
}

func TestSelect(t *testing.T) {
	source := "tcp://127.0.0.1:9001?debug=true&username=root&password=123456&database=testhi"
	connect, err := sqlx.Connect("clickhouse", source)
	if err != nil {
		fmt.Printf("clickhouse open err %s", err.Error())
	}
	defer func() {
		_ = connect.Close()
	}()

	var items []MenuItem
	if err := connect.Select(&items, "select * from default.menu_item limit 10"); err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Printf("id: %d, name: %s, age: %d\n", item.Id, item.Name, item.Age)
	}
}
