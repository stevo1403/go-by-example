package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/stevo1403/go-by-example/apps/comment"
	"github.com/stevo1403/go-by-example/apps/post"
	"github.com/stevo1403/go-by-example/apps/user"
)

func LoadViews() {
	r := gin.Default()

	r.GET("/users", user.GetUsers)
	r.POST("/users", user.CreateUser)
	r.GET("/users/:id", user.GetUser)
	r.PUT("/users/:id", user.UpdateUser)
	r.DELETE("/users/:id", user.DeleteUser)

	r.GET("/comments", comment.ListComments)
	r.POST("/comments/", comment.CreateComment)
	r.GET("/comments/:id", comment.GetComment)
	r.PUT("/comments/:id", comment.UpdateComment)
	r.PATCH("/comments/:id/upvote", comment.UpvoteComment)
	r.PATCH("/comments/:id/downvote", comment.DownVoteComment)
	r.DELETE("/comments/:id", comment.DeleteComment)

	r.GET("/posts", post.ListPosts)
	r.POST("/posts", post.CreatePost)
	r.GET("/posts/:id", post.GetPost)
	r.PUT("/posts/:id", post.UpdatePost)
	r.DELETE("/posts/:id", post.DeletePost)

	r.Run()
}
