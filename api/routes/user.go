package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hedonicadapter/gopher/services/user"
)

func UserRoutes(rg *gin.RouterGroup, userService user.UserService) *gin.RouterGroup {
	rg.GET("users/:id", func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		user, err := userService.Get(id)
		if err != nil {
			ctx.JSONP(500, gin.H{
				"status": "FAILED",
				"data":   user,
			})
		}

		ctx.JSONP(200, gin.H{
			"status": "OK",
			"data":   user,
		})
	})

	return rg
}
