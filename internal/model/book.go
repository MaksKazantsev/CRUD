package model

type Book struct {
	ID     int    `json:"ID"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
