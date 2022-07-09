package models

type Book struct {
	id     int    `json:"ID"`
	author string `json:"author"`
	title  string `json:"title"`
}
