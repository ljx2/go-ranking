package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// r.GET("/hello", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "hello world!")
	// })
	user := r.Group("/user")
	{
		user.POST("/list", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user list")
			ctx.JSON(http.StatusOK, "user list")
		})

		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user add")
		})

		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user delete")
		})
	}
	return r
}
