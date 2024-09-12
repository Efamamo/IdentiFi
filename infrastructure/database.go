package infrastructure

import (
	"fmt"
	"log"

	"github.com/Efamamo/WonderBeam/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "efa"
	password = "mini555"
	dbname   = "wonderbeam"
)

var DB *gorm.DB

func ConnectToDB() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	// Your database logic here
}

func Migrate() {
	DB.AutoMigrate(&domain.User{})
	DB.AutoMigrate(&domain.Location{}, &domain.Lodging{})
	DB.AutoMigrate(&domain.Activity{})
}
