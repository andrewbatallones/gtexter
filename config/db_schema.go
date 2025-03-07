package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	fmt.Println("Database setup")

	fmt.Printf("Using Database URL: %s", os.Getenv("DB_URL"))

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	create_chats(conn)
	create_users(conn)
}

func create_chats(conn *pgx.Conn) {
	_, err := conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS chats (
		name varchar(255)
	);`)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create 'chats' table: %v\n", err)
	}
}

func create_users(conn *pgx.Conn) {
	_, err := conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS users (
		name varchar(255)
	);`)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create 'chats' table: %v\n", err)
	}
}
