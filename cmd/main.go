package main

import (
	"intro-npo-eschelon/pkg/seacher"
	"intro-npo-eschelon/configs"

	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.Use(contentTypeMiddleware)

	router.POST("/checkText", seacher.SearchHandler)

	config.Logger.Fatal(http.ListenAndServe(":"+config.SEACHERPORT, router))
}

func contentTypeMiddleware(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.Next()
}
