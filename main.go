package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

func index(w http.ResponseWriter, r *http.Request) {
	test := "Testing"

	t, _ := template.ParseFiles("app/views/index.html")
	t.Execute(w, test)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"healthcheck\": %v}", testDbConnection())
}

func testDbConnection() bool {
	// Test connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return false
	}
	defer conn.Close(context.Background())

	var test string
	err = conn.QueryRow(context.Background(), "SELECT 'test'").Scan(&test)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false
	}

	return true
}

func main() {
	fmt.Println("server is up...")
	http.HandleFunc("/", index)
	http.HandleFunc("/healthcheck", healthcheck)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
