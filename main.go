package main

import (
	"github.com/argilapp/core/db"
	"github.com/argilapp/core/models/user"
	routes "github.com/argilapp/core/router"
)

func main() {
	db.Connect()
	db.Session.AutoMigrate(&user.User{}) // would like this moved to db package but need to fix the import cycle for that

	// db.Session.Create(&user.User{
	// 	Username:     "test1",
	// 	DisplayName:  "Test User",
	// 	Email:        "test@argil.app",
	// 	PasswordHash: "notarealhash",
	// })

	// db.Session.Create(&user.User{
	// 	Username:     "test2",
	// 	DisplayName:  "Test User",
	// 	Email:        "test2@argil.app",
	// 	PasswordHash: "notarealhash",
	// })

	// db.Session.Create(&user.User{
	// 	Username:     "test2",
	// 	DisplayName:  "Test User",
	// 	Email:        "test3@argil.app",
	// 	PasswordHash: "notarealhash",
	// })

	r := routes.SetupRouter()
	r.Run("0.0.0.0:8081")
}
