package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Session *gorm.DB

func Connect() {
	// this is all temporary crap that shouldntbe here
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "argil_",
			SingularTable: false,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	Session = db
}
