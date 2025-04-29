package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hedonicadapter/gopher/models"
	"github.com/hedonicadapter/gopher/services/queue"
	"github.com/hedonicadapter/gopher/services/user"
)

func UserRoutes(rg *gin.RouterGroup, users user.UserService, queue queue.QueueService) *gin.RouterGroup {
	rg.GET("users/:id", func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		user, err := users.GetById(id)
		if err != nil {
			// TODO: shouldnt be 500
			ctx.JSONP(500, gin.H{
				"status": "FAILED",
				"data":   user,
			})
			return
		}

		// TODO: transaction
		errr := queue.Enqueue(ctx, models.Task{Action: id})
		if errr != nil {
			// TODO: shouldnt be 500
			ctx.JSONP(500, gin.H{
				"status": "FAILED",
				"data":   errr.Error(),
			})
			return
		}

		ctx.JSONP(200, gin.H{
			"status": "OK",
			"data":   user,
		})
	})

	return rg
}
