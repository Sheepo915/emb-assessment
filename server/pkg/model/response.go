package model

type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Locale string `json:"locale"`
	Seed   string `json:"seed"`
	Total  int    `json:"total"`
	Data   []Book `json:"data"`
}
