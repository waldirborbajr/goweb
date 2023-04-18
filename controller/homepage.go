package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/", index)
	r.GET("/contact", contact)
	r.GET("/about", about)
	r.GET("/login", login)
}

func index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/index.html",
		gin.H{
			"title": "Geeksbeginner",
		},
	)
}

func about(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/about.html",
		gin.H{},
	)
}

func contact(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/contact.html",
		gin.H{
			"title": "Contact",
		},
	)
}

func login(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/login.html",
		gin.H{
			"title": "Contact",
		},
	)
}
