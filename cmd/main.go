package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Entry struct {
	ID        int
	Generated int
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Abdu0811", "project")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE large_dataset (id SERIAL PRIMARY KEY,generated INT NOT NULL);")
	if err!=nil{
		log.Fatal(err)
	}
	_,err=db.Exec(`INSERT INTO large_dataset (generated)
	SELECT generate_series(1, 10000000);`)
	if err!=nil{
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id, generated FROM large_dataset")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var entry Entry
		if err := rows.Scan(&entry.ID, &entry.Generated); err != nil {
			log.Fatal(err)
		}
		fmt.Println(entry)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
