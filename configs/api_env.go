package configs

import(
	
	"os"
	"log"
	"github.com/joho/godotenv"
)


func EnvMongoURL() string{

	err := godotenv.Load()

	if err != nil {
		log.Fatal("can't load the .env file")
	}


	return os.Getenv("MONGOURI")
}