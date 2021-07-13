package api

//nesto
// The sql go library is needed to interact with the database
import (
	"encoding/json"
	"fmt"
	repository "github.com/Vukeezy/main/repository"
	"net/http"
)

func GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	/*
		The list of birds is now taken from the store instead of the package level variable we had earlier
	*/

	comment,err := repository.GetStore().GetExercises()

	// Everything else is the same as before
	birdListBytes, err := json.Marshal(comment)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(birdListBytes)
}