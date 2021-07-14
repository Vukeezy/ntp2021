package models

// Weekday - Custom type to hold value for weekday ranging from 1-7
type preparednessLevel int

type muscle int

type exerciseType int



// Declare related constants for each weekday starting with index 1
const (
	HIGH preparednessLevel = iota
	MEDIUM
	LOW
	NONE
)

const (
	BICEPS muscle = iota
	TRICEPS
	QUADRICEPS
	CHEST
	BACK
)

func IntToMuscle(s int) string {
	return [...]string{"Biceps", "Triceps", "Quadriceps", "Chest", "Back"}[s]
}

func IntToPreparednessLevel(s int) string {
	return [...]string{"High", "Medium", "Low", "None"}[s]
}

func IntToExerciseType(s int) string {
	return [...]string{"Weight loss", "Mass gain"}[s]
}



