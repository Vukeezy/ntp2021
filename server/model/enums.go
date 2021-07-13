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

const (
	WEIGHT_LOSS exerciseType = iota
	MASS_GAIN
)
