package models

type Person struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Gender      string `json:"gender"`
	CoolLevel   bool   `json:"coollevel"`
	Name        string `json:"name"`
}
