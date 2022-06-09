package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/zakirkun/microservices-eco/auth/lib"
)

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func (config configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	lib.PanicIfNeed(err)
	return configImpl{}
}
