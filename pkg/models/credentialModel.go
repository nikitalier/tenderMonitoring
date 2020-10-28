package models

import "github.com/dgrijalva/jwt-go"

//Claims ...
type Claims struct {
	Credential Credential
	jwt.StandardClaims
}

//Credential ...
type Credential struct {
	ID          int      `json:"id"`
	Login       string   `json:"login"`
	FullName    string   `json:"full_name" db:"full_name"`
	Authorities []string `json:"authorities"`
}

//JwtToken ...
type JwtToken struct {
	AccessToken string `json:"accessToken"`
}
