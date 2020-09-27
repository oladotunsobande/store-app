package secrets

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

// Secrets Struct for all environment variables
type Secrets struct {
	Environment      string
	DatabaseType     string
	DatabaseProvider string
	ApplicationPort  string
	ApplicationName  string
	DatabaseURL      string
}

func init() {
	//_, b, _, _ := runtime.Caller(0)
	//BasePath := path.Dir(b)
	log.Println(path.Join(".env"))
	if err := godotenv.Load(path.Join(".env")); err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Environment variables has been successfully loaded.")
}

// GetSecrets Get all loaded secrets
func GetSecrets() Secrets {
	var secrets Secrets

	secrets.Environment = os.Getenv("GO_ENV")
	secrets.DatabaseType = os.Getenv("DB_TYPE")
	secrets.ApplicationPort = os.Getenv("APP_PORT")
	secrets.DatabaseProvider = os.Getenv("DB_PROVIDER")
	secrets.ApplicationName = os.Getenv("APP_NAME")
	secrets.DatabaseURL = os.Getenv("DB_URI")

	return secrets
}
