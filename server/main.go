package main

import (
	"database/sql"
	"fmt"
	handler "github.com/Vukeezy/main/api"
	repository "github.com/Vukeezy/main/repository"

	_ "github.com/lib/pq"
	"net/http"

	"github.com/gorilla/mux"
)
// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/comment", handler.GetCommentHandler).Methods("GET")

	//r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	//r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func main() {
	fmt.Println("Starting server...")
	connString := "dbname=postgres user=postgres password=vukasin123 sslmode=disable"
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
	r := newRouter()
	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", r)
}

