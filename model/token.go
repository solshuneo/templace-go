package model

type Token struct {
	Access  string "json:\"access\""
	Refresh string "json:\"refresh\""
}

func NewToken() *Token {
	return &Token{
		Access:  "accessToken",
		Refresh: "refreshToken",
	}
}
