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
			"msg": "post created",
			"ok":  true,
		})
	})
	g.GET("/posts", func(ctx *gin.Context) {
		posts, err := service.GetPosts()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err,
				"ok":    false,
			})
			return
		}
		ctx.JSON(200, gin.H{
			"posts": posts,
			"ok":    true,
		})
	})
	g.GET("/posts/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		post, err := service.GetOnePost(id)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
				"ok":    false,
			})
			return
		}
		if post == nil {
			ctx.JSON(404, gin.H{
				"error": "post not found",
				"ok":    false,
			})
			return
		}

		ctx.JSON(200, gin.H{
			"post": post,
			"ok":   true,
		})
	})
}
