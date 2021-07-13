package models

type Exercise struct {
	Id                    int
	RequestedPreparedness int
	Equipment             bool
	Muscles               []int
	Name                  string
	Description           string
	Type                  int
	Comments              []Comment
}
