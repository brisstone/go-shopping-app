package routes

import (
    "shopping-app/controllers" //Import Controller
    "github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    //All routes related to users comes here
	router.POST("/user", controllers.CreateUser()) //add
	router.GET("/user/:userId", controllers.GetAUser()) //add t
	router.PUT("/user/:userId", controllers.EditAUser()) //add thi
	router.DELETE("/user/:userId", controllers.DeleteAUser()) //a
	router.GET("/users", controllers.GetAllUsers()) //add th
	router.POST("/login", controllers.LoginUser())
	router.POST("/register", controllers.RegisterUser())


	
}