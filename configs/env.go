package configs
// a helper function to load the environment variable 
// using the github.com/joho/godotenv library installed
import (
    // "log"
    "os"
    // "github.com/joho/godotenv"
)

// EnvMongoURI function that checks if the environment variable is correctly loaded 
// and returns the environment variable.

func EnvMongoURI() string {
    // err := godotenv.Load()
    // if err != nil {
    //     log.Fatal("Error loading .env file", err)
    // }

    return os.Getenv("MONGOURI")
}