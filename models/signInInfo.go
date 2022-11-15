package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type SignInInfo struct {
}

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	User_type  string
	jwt.StandardClaims
}
