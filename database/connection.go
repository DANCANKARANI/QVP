package database

import (
    "fmt"
    "log"
    //"os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

// ConnectDB connects to the MySQL database and returns the GORM DB object.
func ConnectDB() *gorm.DB {
    // Load .env file
    log.Println("Welcome")
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Retrieve environment variables
    dbUser := "u3e6lqcmk6bomlmv"
    password := "fL3O93uiXU35Olm7ybP"
    dbName := "booi0hsjshh69jxkimhr"
    host := "booi0hsjshh69jxkimhr-mysql.services.clever-cloud.com"
    port := "21345"
    log.Println(password,dbUser,dbName,host,port)
    // Check if any of the environment variables are empty
    if dbUser == ""|| dbName == ""||password=="" || host == "" || port == "" {
        log.Fatalf("Database configuration variables are missing. Check your .env file.")
    }
   
    // Construct DSN (Data Source Name)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    dbUser, password, host, port, dbName)
    // Connect to the database using GORM v2
    db, err := gorm.Open(mysql.Open(dsn))
    if err != nil {
        log.Fatalf("Failed to connect to the database:n"+err.Error())
    }

    fmt.Println("Database connection established successfully.")
    return db
}
