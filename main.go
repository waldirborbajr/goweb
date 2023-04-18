package main

import (
	"localhost/goweb/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/scss", "./static/scss")
	r.Static("/vendor", "./static/vendor")
	r.Static("/js", "./static/js")
	r.StaticFile("/favicon.ico", "./img/favicon.ico")

	r.LoadHTMLGlob("templates/**/*")
	controller.Router(r)

	log.Println("Server started")

	if err := r.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
