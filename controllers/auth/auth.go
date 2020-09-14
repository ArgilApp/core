package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get the current logged in user
func Authenticate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": "nope",
	})
}
