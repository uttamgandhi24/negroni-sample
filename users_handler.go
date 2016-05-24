package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

func UsersCreateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "posts create")
}

func UserShowHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["userID"]
	fmt.Fprintln(rw, "showing post", id)

	user := context.Get(r, "user")
	fmt.Fprintf(rw, "This is an authenticated request")
	fmt.Fprintf(rw, "Claim content:\n")
	for k, v := range user.(*jwt.Token).Claims {
		fmt.Fprintf(rw, "%s :\t%#v\n", k, v)
	}
}
