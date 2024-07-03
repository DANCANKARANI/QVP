package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

// ConnectDB connects to the MySQL database and returns the GORM DB object.
func ConnectDB() *gorm.DB {
    // Load .env file
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Retrieve environment variables
    userName := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")

    // Check if any of the environment variables are empty
    if userName == "" || password == "" || dbName == "" || host == "" || port == "" {
        log.Fatalf("Database configuration variables are missing. Check your .env file.")
    }

    // Construct DSN (Data Source Name)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        userName, password, host, port, dbName)

    // Connect to the database using GORM v2
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    fmt.Println("Database connection established successfully.")
    return db
}
