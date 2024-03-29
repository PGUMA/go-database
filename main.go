package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"github.com/PGUMA/go-database/ent"
	"github.com/PGUMA/go-database/ent/sampletable"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "host=127.0.0.1 port=15432 user=postgres password=password dbname=dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = db.PingContext(ctx) // dbへの実際に接続を確立しに行かせることで以降のSQL発行時に接続エラーが起こる可能性を軽減
	if err != nil {
		log.Fatal(err)
	}

	it, _ := time.ParseDuration("5s")
	lt, _ := time.ParseDuration("60s")

	db.SetConnMaxIdleTime(it)
	db.SetConnMaxLifetime(lt)
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(2)
}

func main() {

	embededSample()

	fmt.Println("============================================")

	entSample()
}

func embededSample() {
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM sample_table WHERE id < $1`, 4)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int8
		var name string
		var age int8
		var assets int64
		var height float32
		var alive bool
		var birthday string
		var birthAt string
		var note string

		err = rows.Scan(&id, &name, &age, &assets, &height, &alive, &birthday, &birthAt, &note)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v \n", id, name, age, assets, height, alive, birthday, birthAt, note)
	}
}

func entSample() {
	client, err := ent.Open(dialect.Postgres, "host=127.0.0.1 port=15432 user=postgres password=password dbname=dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// Run the auto migration tool.
	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	rec, err := client.Debug().SampleTable.Query().Where(sampletable.IDLT(4)).All(context.Background())
	if err != nil {
		log.Fatalf("failed query: %v", err)
	}

	for _, e := range rec {
		fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v \n", e.ID, e.VarcharColumn, e.IntegerColumn, e.BigintColumn, e.NumericColumn, e.BooleanColumn, e.DateColumn, e.TimestampColumn, e.TextColumn)
	}
}
