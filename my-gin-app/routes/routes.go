package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/username/my-gin-app/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	// ユーザー関連のルート
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
	}

	// 投稿関連のルート
	postGroup := router.Group("/posts")
	{
		postGroup.GET("/", controllers.ListPosts)
		postGroup.GET("/:id", controllers.ShowPost)
		postGroup.POST("/", controllers.CreatePost)
		postGroup.PUT("/:id", controllers.UpdatePost)
		postGroup.DELETE("/:id", controllers.DeletePost)
	}
}
