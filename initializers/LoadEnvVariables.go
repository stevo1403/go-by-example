package initializers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)

	}
}

func LoadDB() {
	LoadEnvVariables()
	var db_directory = os.Getenv("SQLITE_DB_DIRECTORY")
	var db_filename = os.Getenv("SQLITE_DB_NAME")
	var database_file_path = filepath.Join(db_directory, db_filename)

	_, path_err := os.Stat(db_directory)

	if os.IsNotExist(path_err) {
		os.MkdirAll(db_directory, os.ModePerm)
	}

	var dbHandle = sqlite.Open(database_file_path)
	var err error

	DB, err = gorm.Open(dbHandle, &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open '%s': %s", database_file_path, err)
	}

	if err != nil {
		log.Fatalf("Failed to open '%s': %s", database_file_path, err)
	}

}
