package api

//nesto
// The sql go library is needed to interact with the database
import (
	"encoding/json"
	"fmt"
	model "github.com/Vukeezy/main/model"
	"github.com/Vukeezy/main/repository"

	"net/http"
)

//domain/exercises
func GetExercisesHandler(w http.ResponseWriter, r *http.Request) {
	/*
		The list of birds is now taken from the store instead of the package level variable we had earlier
	*/

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