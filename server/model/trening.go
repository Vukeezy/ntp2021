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
}
