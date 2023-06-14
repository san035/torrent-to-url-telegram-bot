package config

import "github.com/joho/godotenv"

var err error

func init() {
	err = godotenv.Load()
}

func Init() error {
	return err
}
