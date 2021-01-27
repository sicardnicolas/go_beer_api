package main

import (
	"go_beer_api/beer_infrastructure"

	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

var routes = []route{
	route{"AllBeers", "GET", "/beers", beer_infrastructure.AllBeers},
	route{"GetBeer", "GET", "/beers/:id", beer_infrastructure.GetBeer},
	route{"AddBeer", "PUT", "/beers", beer_infrastructure.AddBeer},
	route{"SearchBeer", "POST", "/beers", beer_infrastructure.SearchBeer},
}

func main() {
	router := gin.Default()

	router.GET("/login", gin.BasicAuth(gin.Accounts{"toto": "qwerty"}), security.GenerateToken)
	router.HEAD("/verify", security.VerifyToken)

	for _, r := range routes {
		router.Handle(r.Method, r.Pattern, security.VerifyToken, r.HandlerFunc)
	}
	router.Run(":8080")
}
