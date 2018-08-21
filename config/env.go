package config

import (
	"os"
)

func SetEnv() {
	os.Setenv("PORT", "3000")
	os.Setenv("MONGO_HOST", "127.0.0.1")
	os.Setenv("MONGO_PORT", "27017")
	os.Setenv("MONGO_DB", "goauth")
}
