package auth

import (
	"net/http"

	"github.com/argilapp/core/models/user"
	"github.com/gin-gonic/gin"
)

type authenticateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Get the current logged in user
func Authenticate(c *gin.Context) {
	var request authenticateRequest
	c.ShouldBindJSON(&request)

	success := user.Authenticate(request.Username, request.Password)

	if success {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"token":   "notarealtoken",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Invalid username and/or password",
		})
	}
}
