package main

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// database postgresql configurations and connection
	connectionString := "postgresql://posts:12345@db:5432/posts"
	_, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}
	// api configurations and route /
	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	g.Run(":3000")
}
