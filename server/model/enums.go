package models

func IntToMuscle(s int) string {
	return [...]string{"Biceps", "Triceps", "Quadriceps", "Chest", "W"}[s]
}

func IntToPreparednessLevel(s int) string {
	return [...]string{"High", "Medium", "Low", "None"}[s]
}

func IntToExerciseType(s int) string {
	return [...]string{"Weight loss", "Mass gain"}[s]
}

func IntToAgeRecommendation(s int) string {
	return [...]string{"Young", "Medium", "Older"}[s]
}




