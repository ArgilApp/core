package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"

	"github.com/argilapp/core/models/user"
)

//Get the current logged in user
func GetUser(c *gin.Context) {
	currentUser, _ := user.GetUserByID(1) // obviously not the right value

	currentUserSelfView := &user.UserSelfView{}
	deepcopier.Copy(currentUser).To(currentUserSelfView)

	c.JSON(http.StatusOK, currentUserSelfView)
}

func GetUserByUsername(c *gin.Context) {
	username := c.Params.ByName("username")

	searchedUser, err := user.GetUserByUsername(username)

	if err == nil {
		searchedUserProfileView := &user.UserProfileView{}
		err = deepcopier.Copy(searchedUser).To(searchedUserProfileView)

		c.JSON(http.StatusOK, searchedUserProfileView)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
}
