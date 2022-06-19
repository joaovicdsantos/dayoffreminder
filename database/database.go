package database

import (
	"fmt"
	"io/ioutil"
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
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require Timezone=America/Recife",
		dbHost, dbUser, dbPass, dbBase, dbPort,
	)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect Database")
	}
	log.Println("Connection Opened to Database")
}

func Migrate() {
	DBConn.AutoMigrate(&model.DayOff{})
	loadSQLFiles()
	log.Println("Migrated")
}

func loadConfig() {
	godotenv.Load()
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbBase = os.Getenv("DB_BASE")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
}

func loadSQLFiles() {
	files := loadMigrationFiles()
	for _, file := range files {
		sql, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		DBConn.Exec(string(sql))
	}
}

func loadMigrationFiles() []string {
	currentDirectory, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/database/migrations/", currentDirectory)

	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var formatedFiles []string
	for _, file := range files {
		formatedFiles = append(formatedFiles, fmt.Sprintf("%s/%s", filePath, file.Name()))
	}

	return formatedFiles
}
