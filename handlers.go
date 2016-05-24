package main

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func AddHandlers(router *mux.Router) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com"},
	})

	// authorizations/login
	auth := router.PathPrefix("/authorizations").Subrouter()
	auth.Path("/login").Methods("POST").HandlerFunc(LoginHandler)

	// JWT Middleware
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	// Create userBase with Gorilla router and jwtMiddleware with negroni
	// Add GET and POST /users/{userID}
	userBase := mux.NewRouter()
	users := userBase.PathPrefix("/users/{userID}").Subrouter()
	users.Methods("GET").HandlerFunc(UserShowHandler)
	users.Methods("POST").HandlerFunc(UsersCreateHandler)

	router.PathPrefix("/users").Handler(negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.NewLogger(),
		negroni.Wrap(userBase),
		c,
	))
}
