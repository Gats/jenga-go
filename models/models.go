package models

import(

)

type TokenResponse struct {
	TokenType string `json:"token_type"`
	IssuedAt string `json:"issued_at"`
	ExpiresIn string `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

type ErrorResponse struct {
	Code	string	`json:"code"`
	Error 	string 	`json:"error"`
	Message string 	`json:"message"`
}