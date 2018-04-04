package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username	string 	`json: 	username`
	Password	string	`json:	password`
}

var router *gin.Engine

func main() {

	router = gin.Default()
	initializeRoutes()
	router.Run()
}