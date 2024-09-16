package comment

import (
	"github.com/stevo1403/go-by-example/apps/post"
	"github.com/stevo1403/go-by-example/apps/user"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	AuthorID  int `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign key
	Author    user.User
	PostID    int `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign key
	Post      post.Post
	Body      string
	upvotes   int
	downvotes int
}
