package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

type Storage struct {
	Host     string
	Username string
	Password string
	Database string
	Port     string
	SSLmode  string
	TimeZone string
}

func InitDB() {
	var psqlConf = Storage{
		Host:     os.Getenv("HOST"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Database: os.Getenv("DATABASE"),
		Port:     os.Getenv("PORT_DB"),
		SSLmode:  "disable",
		TimeZone: "Europe/Warsaw",
	}
	ConnectToDB(&psqlConf)
}

func ConnectToDB(conf *Storage) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conf.Host, conf.Username, conf.Password, conf.Database, conf.Port, conf.SSLmode, conf.TimeZone)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}
}
