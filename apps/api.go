package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/stevo1403/go-by-example/apps/comment"
	"github.com/stevo1403/go-by-example/apps/post"
	"github.com/stevo1403/go-by-example/apps/user"
	middlewares "github.com/stevo1403/go-by-example/middlewares"
)

func LoadViews() {
	r := gin.Default()

	user_router := r.Group("/users")
	user_router.Use(middlewares.UserAuthMiddleware)

	user_router.GET("", user.GetUsers)
	user_router.POST("", user.CreateUser)
	user_router.GET("/:id", user.GetUser)
	user_router.PUT("/:id", user.UpdateUser)
	user_router.DELETE("/:id", user.DeleteUser)

	comment_router := r.Group("/comments")
	comment_router.Use(middlewares.UserAuthMiddleware)

	comment_router.GET("", comment.ListComments)
	comment_router.POST("", comment.CreateComment)
	comment_router.GET("/:id", comment.GetComment)
	comment_router.PUT("/:id", comment.UpdateComment)
	comment_router.PATCH("/:id/upvote", comment.UpvoteComment)
	comment_router.PATCH("/:id/downvote", comment.DownVoteComment)
	comment_router.DELETE("/:id", comment.DeleteComment)

	post_router := r.Group("/posts")
	post_router.GET("", post.ListPosts)
	post_router.POST("", post.CreatePost)
	post_router.GET("/:id", post.GetPost)
	post_router.PUT("/:id", post.UpdatePost)
	post_router.DELETE("/:id", post.DeletePost)
	post_router.Use(middlewares.UserAuthMiddleware)

	auth_router := r.Group("/auth")

	auth_router.POST("/login", user.AuthenticateUser)
	auth_router.POST("/signup", user.CreateUser)

	// r.GET("/users", user.GetUsers)
	// r.POST("/users", user.CreateUser)
	// r.GET("/users/:id", user.GetUser)
	// r.PUT("/users/:id", user.UpdateUser)
	// r.DELETE("/users/:id", user.DeleteUser)

	// r.GET("/comments", comment.ListComments)
	// r.POST("/comments/", comment.CreateComment)
	// r.GET("/comments/:id", comment.GetComment)
	// r.PUT("/comments/:id", comment.UpdateComment)
	// r.PATCH("/comments/:id/upvote", comment.UpvoteComment)
	// r.PATCH("/comments/:id/downvote", comment.DownVoteComment)
	// r.DELETE("/comments/:id", comment.DeleteComment)

	// r.GET("/posts", post.ListPosts)
	// r.POST("/posts", post.CreatePost)
	// r.GET("/posts/:id", post.GetPost)
	// r.PUT("/posts/:id", post.UpdatePost)
	// r.DELETE("/posts/:id", post.DeletePost)

	// r.POST("/auth/login", user.AuthenticateUser)

	r.Run()
}
