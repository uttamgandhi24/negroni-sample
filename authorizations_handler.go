package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

/*	Auth handlers	*/
func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "login")

	authHeader := r.Header.Get("Authorization")
	if len(authHeader) == 0 {
		rw.WriteHeader(http.StatusUnauthorized)
		fmt.Errorf("Authorization header missing")
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "basic" {
		rw.WriteHeader(http.StatusUnauthorized)
		fmt.Errorf("Authorization header format must be Basic {token}")
		return
	}

	data, err := base64.StdEncoding.DecodeString(authHeaderParts[1])
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	userpwd := strings.Split(string(data), ":")
	encodedPassword := base64.StdEncoding.EncodeToString([]byte(userpwd[1]))

	// fetch userid with help of username and encoded pwd
	userID := getUserIDByNameAndPassword(userpwd[0], encodedPassword)

	if len(userID) == 0 {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	jwtToken, err := generateJWT(userID)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
  	if err := json.NewEncoder(rw).Encode(jwtToken); err != nil {
    	panic(err)
  	}
}

func generateJWT(userID string) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token.Claims["userID"] = userID
	// Sign and get the complete encoded token as a string
	tokenString, err = token.SignedString([]byte("My Secret"))
	fmt.Println("The token is", tokenString)
	return tokenString, err
}
