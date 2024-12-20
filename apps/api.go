package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/stevo1403/go-by-example/apps/comment"
	"github.com/stevo1403/go-by-example/apps/post"
	"github.com/stevo1403/go-by-example/apps/user"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	"github.com/stevo1403/go-by-example/docs"
	middlewares "github.com/stevo1403/go-by-example/middlewares"
)

func LoadViews() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		user_router := v1.Group("/users")
		{
			user_router.POST("", user.CreateUser)

			user_router.Use(middlewares.UserAuthMiddleware)

			user_router.GET("", user.GetUsers)
			user_router.GET("/:id", user.GetUser)
			user_router.PUT("/:id", user.UpdateUser)
			user_router.DELETE("/:id", user.DeleteUser)
		}

		comment_router := v1.Group("/comments")
		{
			comment_router.GET("/:id", comment.GetComment)
			comment_router.GET("", comment.ListComments)
			comment_router.Use(middlewares.UserAuthMiddleware)

			comment_router.POST("", comment.CreateComment)
			comment_router.PUT("/:id", comment.UpdateComment)
			comment_router.PATCH("/:id/upvote", comment.UpvoteComment)
			comment_router.PATCH("/:id/downvote", comment.DownVoteComment)
			comment_router.DELETE("/:id", comment.DeleteComment)
		}

		post_router := v1.Group("/posts")
		{
			post_router.GET("/:id", post.GetPost)
			post_router.GET("", post.ListPosts)
			post_router.Use(middlewares.UserAuthMiddleware)

			post_router.POST("", post.CreatePost)
			post_router.PUT("/:id", post.UpdatePost)
			post_router.DELETE("/:id", post.DeletePost)
		}

		auth_router := v1.Group("/auth")
		{
			auth_router.POST("/login", user.AuthenticateUser)
			auth_router.POST("/signup", user.CreateUser)
		}
	}
	
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("localhost:8080")
}
