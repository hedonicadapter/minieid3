package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hedonicadapter/gopher/api/routes"
	"github.com/hedonicadapter/gopher/config"
	"github.com/hedonicadapter/gopher/models"
	"github.com/hedonicadapter/gopher/services/queue"
	"github.com/hedonicadapter/gopher/services/user"
)

func main() {
	config.InitEnv()
	db := config.InitDb()

	rdb := config.InitRedis()
	config.IdempotentDummyData(db)

	r := gin.Default()

	userService := user.InitService(db)
	queueService := queue.InitService(rdb, "main")
	routes.UserRoutes(r.Group("api"), userService, queueService)

	r.GET("health", func(ctx *gin.Context) {
		ctx.JSONP(200, gin.H{
			"status": "OK",
		})
	})

	go queueService.Poll(context.Background(), func(task models.Task) any {
		fmt.Println(task.Action)

		return ""
	})

	r.Run()
}
