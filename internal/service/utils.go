package service

import (
	"fmt"

	"github.com/itsDrac/godo/utils"
)

func HashPassword(password string) (string, error) {
	cost := utils.GetEnvAsInt("HASH_COST", 12)
	return "", fmt.Errorf("not implemented %d", cost)
}

func VerifyPassword(password, hash string) bool {
	return false
}