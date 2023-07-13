package api

import (
	"errors"
	"os"
)

// check service admin
func CheckServiceAdmin() error {
	serviceAdmin := os.Getenv("SERVICE_ADMIN")
	if serviceAdmin == "" {
		return errors.New("Service admin is empty")
	}
	return nil
}
