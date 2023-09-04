package model

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB, err = gorm.Open(sqlite.Open("db.sqlite"))
)

type Todo struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Todo      string
	Completed bool
}

func InitializeDB() error {
	if err != nil {
		_, err := os.Create("db.sqlite")
		DB, err = gorm.Open(sqlite.Open("db.sqlite"))
		if err != nil {
			return err
		}
	}

	err = DB.AutoMigrate(&Todo{})

	return err
}
