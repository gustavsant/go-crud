package security

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func RetrieveSecurityToken() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	return os.Getenv("SECRET_KEYWORD")

}
