package routing

import (
	handler "github.com/Vukeezy/main/api"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/exercises", handler.GetExercisesHandler).Methods("GET")
	r.HandleFunc("/getExercise", handler.GetExerciseById).Methods("GET")
	r.HandleFunc("/rateExercise", handler.RateExercise).Methods("POST")
	r.HandleFunc("/rateComment", handler.RateComment).Methods("POST")
	r.HandleFunc("/commentExercise", handler.AddComment).Methods("POST")
	r.HandleFunc("/searchExercises", handler.SearchExercises).Methods("POST")
	r.HandleFunc("/getExercisesSorted", handler.GetExerciseSortedHandler).Methods("GET")
	r.HandleFunc("/getPersonalExercises", handler.GetPersonalExercisess).Methods("POST")


	return r
}