package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/argilapp/core/controllers/auth"
	"github.com/argilapp/core/controllers/user"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("version", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"branch":  gin.Mode(),
				"version": "unknown",
			})
		})

		userGroup := api.Group("/user")
		{
			userGroup.GET("/", user.GetUser)
			subUserGroup := userGroup.Group("/:username")
			{
				subUserGroup.GET("", user.GetUserByUsername)
				subUserGroup.GET("/repositories", NotImplemented)
			}
		}

		authGroup := api.Group("/auth")
		{
			authGroup.POST("", auth.Authenticate)
			authGroup.DELETE("", NotImplemented)
		}
	}

	return r
}

// Holding not implemented yet controller
func NotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not Implemented Yet",
	})
}
