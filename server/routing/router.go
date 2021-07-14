package routing

import (
	handler "github.com/Vukeezy/main/api"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/exercises", handler.GetExercisesHandler).Methods("GET")
	r.HandleFunc("/rateExercise", handler.RateExercise).Methods("POST")
	r.HandleFunc("/rateComment", handler.RateComment).Methods("POST")


	//r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	//r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}