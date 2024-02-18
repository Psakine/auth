package config

import "github.com/joho/godotenv"

// Load loads values from env
func Load(path string) error {
	err := godotenv.Load(path)

	if err != nil {
		return err
	}

	return err
}
