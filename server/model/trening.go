package models

type Exercise struct {
	Id                    int `json:"id"`
	RequestedPreparedness int `json:"requestedpreparedness"`
	Equipment             bool `json:"equipment"`
	Muscles               []int
	Name                  string `json:"name"`
	Description           string `json:"description"`
	Type                  int `json:"type"`
	Comments              []Comment
	Rates				  []int
}

func GetExerciseDTO (exercise *Exercise) ExerciseDTO {
	var exerciseDTO = ExerciseDTO{Name: exercise.Name, Equipment: exercise.Equipment, Description: exercise.Description,
								Type: IntToExerciseType(exercise.Type), RequestedPreparedness: IntToPreparednessLevel(exercise.RequestedPreparedness),
								Comments: exercise.Comments, Rates: exercise.Rates}
	for i := 0; i < len(exercise.Muscles); i++ {
		exerciseDTO.Muscles = append(exerciseDTO.Muscles, IntToMuscle(exercise.Muscles[i]))
	}
	return exerciseDTO
}

type ExerciseDTO struct {
	RequestedPreparedness string `json:"requestedpreparedness"`
	Equipment             bool `json:"equipment"`
	Muscles               []string
	Name                  string `json:"name"`
	Description           string `json:"description"`
	Type                  string `json:"type"`
	Comments              []Comment
	Rates				  []int
}