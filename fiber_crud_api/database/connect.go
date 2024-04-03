package database

import (
	"fmt"
	"strconv"

	"github.com/wortiz1027/golang/fiber_crud_api/config"
	"github.com/wortiz1027/golang/fiber_crud_api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("Faild to parse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
					config.Config("DB_HOST"), port, config.Config("DB_USERNME"),
				config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Fail to connect to database")
	}

	fmt.Println("Connection is open to database...")

	DB.AutoMigrate(&models.Task{})
}
