package db

import (
	"fmt"
	config "testTask/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func DB(config *config.Config) *gorm.DB {
	connString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Pass, config.Database.DBName)
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
