package main

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/database"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/post"
	"github.com/gin-gonic/gin"
)

func main() {

	// database postgresql configurations and connection
	connectionString := "postgresql://posts:1234@db:5432/posts"
	conn, err := database.NewConnection(connectionString)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer conn.Close()
	// setup repository
	repository := post.NewRepository(conn)
	service := post.NewService(repository)

	// api configurations and route /
	g := gin.Default()
	http.SetupRoutes(g, service)
	g.Run(":3000")

}
