package main

import (
	"authentication-service/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const webPort = "80"

var counts uint16 = 0

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	conn := connectToDB()
	if conn == nil {
		log.Panic("Could not connect to Postgres")
	}

	app := &Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Println("Error connecting to DB:", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	// dsn := "host=postgres port=54322 user=postgres password=password dbname=users sslmode=disable timezone=UTC connect_timeout=5"
	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not ready yet ...")
			counts++
		} else {
			log.Println("Connected to Postgres")
			return conn
		}
		if counts > 10 {
			log.Println("Cannot connect to postgres")
			return nil
		}
		time.Sleep(time.Second * 2)
		log.Println("DB not ready. Sleeping for 2 seconds.")
	}
}
