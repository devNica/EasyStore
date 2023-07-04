package configurations

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filesname ...string) Config {
	err := godotenv.Load(filesname...)
	if err != nil {
		panic(err)
	}

	return &configImpl{}
}
