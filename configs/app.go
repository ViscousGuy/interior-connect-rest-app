package configs

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/joho/godotenv"         // For loading .env
)

func init() {
	// Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, some features might not work")
	}

	// Database connection details
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	// Register the MySQL driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Register the database connection
	err = orm.RegisterDataBase("default", "mysql", connStr)
	if err != nil {
		panic(fmt.Errorf("error registering database: %s", err))
	}
	// Test Connection
	o := orm.NewOrm()
	if err := o.Using("default"); err != nil {
		panic(fmt.Errorf("failed to connect to the database: %s", err))
	} else {
		fmt.Println("Database connection established successfully!")
	}
	orm.Debug = false
}
