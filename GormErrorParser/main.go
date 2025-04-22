package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	//"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"unique"`
	Phone string `gorm:"unique"`
}

var (
	DuplicateErrorCode = "23505"
)

type ErrorDetails struct {
	Message    string // Full error message
	SQLState   string // SQLSTATE code (e.g., "23505" for unique constraint violation)
	Constraint string // Name of the constraint that caused the error
}

// GormErrorParser extracts structured details from a given GORM error.
// It identifies SQLSTATE codes and constraint names from the error message.
func GormErrorParser(err error) ErrorDetails {
	// Ensure the error is not nil to avoid dereferencing issues.
	if err == nil {
		return ErrorDetails{}
	}

	// Initialize default values.
	var sqlstate, constraint string
	errorMessage := err.Error()

	// Check if the error is a PostgreSQL unique constraint violation.
	if strings.Contains(errorMessage, DuplicateErrorCode) {
		sqlstate = DuplicateErrorCode
	}

	// Extract the constraint name from the error message if it exists.
	startIndex := strings.Index(errorMessage, "\"")
	endIndex := strings.LastIndex(errorMessage, "\"")
	if startIndex != -1 && endIndex != -1 && startIndex < endIndex {
		constraint = errorMessage[startIndex+1 : endIndex]
	}

	// Return the structured error details.
	return ErrorDetails{
		Message:    errorMessage,
		SQLState:   sqlstate,
		Constraint: constraint,
	}
}

func main() {
	dsn := "<your_Postgres_Dsn>"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the schema
	db.AutoMigrate(&User{})

	// Insert first record (Success)
	// user1 := User{Email: "test@example.com", Phone: "8899009988"}
	// if err := db.Create(&user1).Error; err != nil {
	// 	log.Fatalf("Error inserting user: %v", err)
	// }

	// Insert duplicate record (Should fail)
	user2 := User{Email: "test@example.com", Phone: "8899009988"}
	tx := db.WithContext(context.Background()).Begin()
	err = tx.Create(&user2).Error

	if err != nil {
		// // Check if the error message contains the unique constraint violation message
		// if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		// 	fmt.Println("✅ Unique constraint violation detected!")
		// } else {
		// 	fmt.Println("❌ Some other error:", err)
		// }
		fmt.Println(GormErrorParser(err).Constraint)
	} else {
		fmt.Println("Inserted successfully")
	}
}
