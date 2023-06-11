package main

import (
    "shopping-app/configs" //add this
    "github.com/gin-gonic/gin"
	"shopping-app/routes" 
)

func main() {
		// Initialize a Gin router using the Default configuration. 
		// The Default function configures Gin router with default middlewares (logger and recovery).
        router := gin.Default()

		 //run database
		 configs.ConnectDB()

		//routes
		routes.UserRoute(router) //add th

		// route to / path and a handler function that returns a JSON of Hello from Gin-gonic & mongoDB.
        router.GET("/", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "data": "Hello from My Go Shopping App (Gin-gonic & mongoDB)",
                })
        })

		//Configure port
        // router.Run("localhost:4000") 
		router.Run(":80")
}