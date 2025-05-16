package main

import (
	"fmt"

	server "github.com/TusharKM1224/Server"
	types "github.com/TusharKM1224/Types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	postgresDSN := "postgresql://postgres:root@localhost:5432/notificationdb"
	db, err := connectToDb(postgresDSN)
	if err != nil {
		fmt.Print("Error : %w", err)
	}
	handlerInstance := server.Initiateserver(db)

}

func connectToDb(Dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&types.Notification{}); err != nil {
		return nil, err
	}
	return db, nil

}
