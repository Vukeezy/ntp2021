package main

import (
	"database/sql"
	"fmt"
	"github.com/Vukeezy/main/repository"
	"github.com/Vukeezy/main/routing"

	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	connString := "dbname=vezbe_db user=postgres password=vukasin123 sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	repository.InitStore(&repository.DbStore{Db: db})

	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := routing.NewRouter()
	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", r)
}

