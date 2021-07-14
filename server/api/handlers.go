package api

//nesto
// The sql go library is needed to interact with the database
import (
	"encoding/json"
	"fmt"
	model "github.com/Vukeezy/main/model"
	"github.com/Vukeezy/main/repository"
	"strconv"

	"net/http"
)

//domain/exercises
func GetExercisesHandler(w http.ResponseWriter, r *http.Request) {
	exercises,err := repository.GetStore().GetExercises()

	var exercisesDTO []model.ExerciseDTO

	for i := 0; i < len(exercises) ; i++ {
		print(i)
		exercisesDTO = append(exercisesDTO, model.GetExerciseDTO(exercises[i]))
	}

	// Everything else is the same as before
	exercisesBytes, err := json.Marshal(exercisesDTO)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(exercisesBytes)

}


//domain/rateExercises
func RateExercise(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	rate, _ := strconv.Atoi(r.Form.Get("rate"))
	exerciseId, _ := strconv.Atoi(r.Form.Get("exerciseId"))

	err := repository.GetStore().RateExercise(exerciseId, rate)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

//domain/rateExercises
func RateComment(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	rate, _ := strconv.Atoi(r.Form.Get("rate"))
	commentId, _ := strconv.Atoi(r.Form.Get("commentId"))

	err := repository.GetStore().RateComment(commentId, rate)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
