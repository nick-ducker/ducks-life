package models

type Rambling struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
}
