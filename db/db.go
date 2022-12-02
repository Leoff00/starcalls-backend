package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/starcalls-backend/db/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func StartDB() {
	errEnv := godotenv.Load("./.env")

	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	var dbUSr, dbPwd = os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PWD")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=PG_Q_DB port=5432", dbUSr, dbPwd)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		panic("failed to connect database, cause ->" + err.Error())
	}

	db = database
	config, _ := db.DB()
	migrations.RunMigration(db)
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	fmt.Println("Connection with Db started!")
}

func GetDatabase() *gorm.DB {
	return db
}

func CloseDB(db *gorm.DB) error {
	exitDB, err := db.DB()

	if err != nil {
		log.Fatal("Failed in opening db", err.Error())
	}

	return exitDB.Close()
}
