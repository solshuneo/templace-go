package model

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Content any    `json:"content"`
}
