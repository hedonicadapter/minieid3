package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	rg.GET("users/:id", func(ctx *gin.Context) {
		id, found := ctx.Params.Get("id")
		if !found {
			fmt.Println("error")
		}

		ctx.JSONP(200, gin.H{
			"status": "OK",
			"data":   id,
		})
	})

	return rg
}
