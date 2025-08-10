package model

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	Image       string `json:"image"`
	Published   string `json:"published"`
	Publisher   string `json:"publisher"`
}
