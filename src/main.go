package main

import (
	"tugas-rest-api/controllers/gamescontroller"
	"tugas-rest-api/controllers/kepemilikancontroller"
	"tugas-rest-api/controllers/userscontroller"
	"tugas-rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Endpoint API

	// 1. Users
	r.GET("/rest/api/users", userscontroller.ShowAll)
	r.GET("/rest/api/users/:id", userscontroller.Show)
	r.POST("/rest/api/user", userscontroller.Create)
	r.PUT("/rest/api/user/:id", userscontroller.Update)
	r.DELETE("/rest/api/user/:id", userscontroller.Delete)

	r.GET("/rest/api/password/:id", userscontroller.ComparePassword) // Untuk mengcompare password

	// 2. Games
	r.GET("/rest/api/games", gamescontroller.ShowAll)
	r.GET("/rest/api/games/:id", gamescontroller.Show)
	r.POST("/rest/api/game", gamescontroller.Create)
	r.PUT("/rest/api/game/:id", gamescontroller.Update)
	r.DELETE("/rest/api/game/:id", gamescontroller.Delete)

	// 3. Kepemilikan
	r.GET("/rest/api/kepemilikan", kepemilikancontroller.ShowAll)
	r.GET("/rest/api/kepemilikan/:id", kepemilikancontroller.Show)
	r.POST("/rest/api/kepemilikan", kepemilikancontroller.Create)
	r.DELETE("/rest/api/kepemilikan/:id", kepemilikancontroller.Delete)

	r.Run()
}
