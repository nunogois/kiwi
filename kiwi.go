package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://github.com/nunogois/kiwi")
	})

	handleRoutes(router)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
