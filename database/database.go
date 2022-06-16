package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joaovicdsantos/dayoffreminder/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	dbHost string
	dbPort string
	dbBase string
	dbUser string
	dbPass string
)

func InitDatabase() {
	var err error

	loadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Recife",
		dbHost, dbUser, dbPass, dbBase, dbPort,
	)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect Database")
	}
	fmt.Println("Connection Opened to Database")
}

func Migrate() {
	DBConn.AutoMigrate(&model.DayOff{})
	fmt.Println("Migrated")
}

func loadConfig() {
	godotenv.Load()
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbBase = os.Getenv("DB_BASE")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
}
