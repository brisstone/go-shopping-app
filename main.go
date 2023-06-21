package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"shopping-app/configs" //add this
	"shopping-app/routes"
	// "time"
)

func main() {

		// Initialize a Gin router using the Default configuration.
	// The Default function configures Gin router with default middlewares (logger and recovery).

	router := gin.Default()


	// CORS for http://localhost:5173 and http://localhost:5173 origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*", "http://localhost:5173"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "http://localhost:5173"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	// router.Run()
	router.Use(cors.Default())



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
	router.Run("localhost:8080")
	// router.Run(":80")
	//Access on localhost:80
}
