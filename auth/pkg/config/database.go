package config

import (
	"fmt"
	"strconv"

	"github.com/zakirkun/microservices-eco/auth/lib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db interface {
	Connect(configuration Config) *gorm.DB
	Close()
}

type database struct {
	db *gorm.DB
}

func NewDatabase(db *gorm.DB) *database {
	return &database{db}
}

func (con *database) Connect(configuration Config) *gorm.DB {
	PORT, _ := strconv.Atoi(configuration.Get("DB_PORT"))
	SSL, _ := strconv.ParseBool(configuration.Get("DB_SSL"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%t", configuration.Get("DB_HOST"), configuration.Get("DB_USER"), configuration.Get("DB_PASS"), configuration.Get("DB_NAME"), PORT, SSL)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	lib.PanicIfNeed(err)

	return db
}
