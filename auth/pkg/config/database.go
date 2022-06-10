package config

import (
	"fmt"
	"strconv"

	"github.com/zakirkun/microservices-eco/auth/lib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(configuration Config) *gorm.DB {
	PORT, _ := strconv.Atoi(configuration.Get("DB_PORT"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", configuration.Get("DB_HOST"), configuration.Get("DB_USER"), configuration.Get("DB_PASS"), configuration.Get("DB_NAME"), PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	lib.PanicIfNeed(err)

	return db
}
