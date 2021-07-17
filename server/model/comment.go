package models

type Comment struct {
	Id       int    `json:"id"`
	FullName string `json:"fullName"`
	Content  string `json:"content"`
	Rates    []int
	Rate	 float64 `json:"rate"`
}

//custom getter for rates
func GetRate() int {
	//TODO: Calculate
	return 5
}
