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
	Age					  int
}

func GetExerciseDTO (exercise *Exercise) ExerciseDTO {
	rate_sum := 0
	var rate_avg float64 = 0.0
	if len(exercise.Rates) > 0 {
		for i := 0; i < len(exercise.Rates); i++{
			rate_sum += exercise.Rates[i]
		}
		rate_avg = float64(rate_sum) / float64(len(exercise.Rates))
	}

	var exerciseDTO = ExerciseDTO{Id: exercise.Id, Name: exercise.Name, Equipment: exercise.Equipment, Description: exercise.Description,
								Type: IntToExerciseType(exercise.Type), RequestedPreparedness: IntToPreparednessLevel(exercise.RequestedPreparedness),
								Comments: exercise.Comments, Rate: rate_avg, Age: IntToAgeRecommendation(exercise.Age)}
	for i := 0; i < len(exercise.Muscles); i++ {
		exerciseDTO.Muscles = append(exerciseDTO.Muscles, IntToMuscle(exercise.Muscles[i]))
	}
	return exerciseDTO
}

type ExerciseDTO struct {
	Id					  int `json:"id"`
	RequestedPreparedness string `json:"requestedpreparedness"`
	Equipment             bool `json:"equipment"`
	Muscles               []string
	Name                  string `json:"name"`
	Description           string `json:"description"`
	Type                  string `json:"type"`
	Comments              []Comment
	Rate				  float64 `json:"rate"`
	Age					  string `json:"ageRecommendation"`
}