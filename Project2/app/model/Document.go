package model

type Document struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Title    string `json:"title"`
}
