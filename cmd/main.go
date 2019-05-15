package main

import (
	"intro-npo-eschelon/configs"
	"intro-npo-eschelon/pkg/seacher"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(contentTypeMiddleware)

	router.POST("/checkText", seacher.SearchHandler)

	configs.Logger.Fatal(http.ListenAndServe(":"+configs.SEACHERPORT, router))
}

func contentTypeMiddleware(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.Next()
}
