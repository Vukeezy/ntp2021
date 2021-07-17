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

//domain/exercises
func GetExerciseSortedHandler(w http.ResponseWriter, r *http.Request) {
	exercises,err := repository.GetStore().GetExercisesSorted()

	var exercisesDTO []model.ExerciseDTO

	for i := 0; i < len(exercises) ; i++ {
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

//domain/rateComment
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

//domain/rateComment
func AddComment(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	exercise_id, _ := strconv.Atoi(r.Form.Get("exerciseId"))
	comment_content := r.Form.Get("content")
	comment_fullname := r.Form.Get("fullname")

	err := repository.GetStore().AddComment(exercise_id, comment_content, comment_fullname)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

//domain/exercises
func SearchExercises(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	requestedPreparedness, _ := strconv.Atoi(r.Form.Get("requestedPreparedness"))
	muscle, _ := strconv.Atoi(r.Form.Get("muscle"))
	equipment, _ := strconv.Atoi(r.Form.Get("equipment"))
	name := r.Form.Get("name")

	exercises,err := repository.GetStore().SearchExercises(requestedPreparedness,equipment,muscle,name)

	var exercisesDTO []model.ExerciseDTO

	for i := 0; i < len(exercises) ; i++ {
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

//domain/exercises
func GetPersonalExercisess(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	age, _ := strconv.Atoi(r.Form.Get("age"))
	preparedness, _ := strconv.Atoi(r.Form.Get("preparedness"))
	goal, _ := strconv.Atoi(r.Form.Get("goal"))
	exercises,err := repository.GetStore().GetPersonalExercises(age,preparedness,goal)

	var exercisesDTO []model.ExerciseDTO

	for i := 0; i < len(exercises) ; i++ {
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

func GetExerciseById(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	requestedPreparedness, _ := strconv.Atoi(r.Form.Get("exerciseId"))

	exercise,err := repository.GetStore().GetExerciseById(requestedPreparedness)

	var exercisesDTO model.ExerciseDTO =  model.GetExerciseDTO(exercise)

	// Everything else is the same as before
	exercisesBytes, err := json.Marshal(exercisesDTO)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(exercisesBytes)

}
