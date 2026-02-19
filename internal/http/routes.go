package http

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/post"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(g *gin.Engine, service *post.Service) {
	g.POST("/posts", func(ctx *gin.Context) {
		var postToCreate post.CreatePostDTO
		if err := ctx.ShouldBindJSON(&postToCreate); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
				"ok":    false,
			})
			return
		}
		if err := service.AddPost(postToCreate); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
				"ok":    false,
			})
			return
		}
		ctx.JSON(201, gin.H{
			"msg": "user created",
			"ok":  true,
		})
	})
}
