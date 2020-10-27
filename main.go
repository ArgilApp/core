package main

import (
	"github.com/argilapp/core/db"
	"github.com/argilapp/core/models/user"
	"github.com/argilapp/core/router"
)

func main() {
	db.Connect()
	db.Session.AutoMigrate(&user.User{}) // would like this moved to db package but need to fix the import cycle for that

	// for i := 1; i <= 5; i++ {
	// 	newUser := &user.User{
	// 		Username:    fmt.Sprintf("test%d", i),
	// 		DisplayName: fmt.Sprintf("Test User (%d)", i),
	// 		Email:       fmt.Sprintf("test%d@argil.app", i),
	// 	}
	// 	newUser.SetPassword(fmt.Sprintf("password%d", i))
	// 	db.Session.Create(newUser)
	// }

	r := router.SetupRouter()
	r.Run("0.0.0.0:8081")
}
