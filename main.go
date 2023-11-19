package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/learning-go-lib/programming"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, welcome to the learning-go-api"})
	})

	base := r.Group("/v1")

	p := programming.ProgrammingFunctions{}
	programming.SetRouterGroup(&p, base)

	r.Run() // listen and serve on 0.0.0.0:8080
}
