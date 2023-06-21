package configs

// a helper function to load the environment variable
// using the github.com/joho/godotenv library installed
import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// EnvMongoURI function that checks if the environment variable is correctly loaded
// and returns the environment variable.

func EnvMongoURI() string {


	envValue := os.Getenv("GO_ENV")
    //Adding below fails on render, so I added GO_ENV in my machine environment variables 
    //by running: <export GO_ENV=development>  in my terminal 
    // To confirm value, run: echo $GO_ENV
    log.Print(os.Getenv("GO_ENV"))

	if envValue == "development" {
        //To prevent it from running on render 
		log.Print(os.Getenv("GO_ENV"))
        // godotenv.Load() function looks for the .env file in the current working directory. 
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	return os.Getenv("MONGOURI")
}
