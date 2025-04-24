package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hedonicadapter/gopher/api/routes"
)

func main() {
	r := gin.Default()

	// add routes
	r.GET("health", func(ctx *gin.Context) {
		ctx.JSONP(200, gin.H{
			"status": "OK",
		})
	})

	routes.UserRoutes(r.Group("api"))

	r.Run()
}
