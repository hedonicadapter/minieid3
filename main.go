package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hedonicadapter/gopher/api/routes"
	"github.com/hedonicadapter/gopher/config"
	"github.com/hedonicadapter/gopher/services/user"
	// "github.com/hedonicadapter/gopher/api/routes"
)

func main() {
	config.InitEnv()

	pool := config.InitDb()
	defer pool.Close()

	r := gin.Default()

	userService := user.InitService(pool)
	routes.UserRoutes(r.Group("api"), userService)

	r.GET("health", func(ctx *gin.Context) {
		ctx.JSONP(200, gin.H{
			"status": "OK",
		})
	})

	r.Run()
}
