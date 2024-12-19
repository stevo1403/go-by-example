package comment

import (
	"github.com/stevo1403/go-by-example/apps/post"
	"github.com/stevo1403/go-by-example/apps/user"
	app "github.com/stevo1403/go-by-example/initializers"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	AuthorID  uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign key
	Author    user.User
	PostID    uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign key
	Post      post.Post
	Body      string
	UpVotes   uint64
	DownVotes uint64
}

func (c Comment) UpdateFields() {
	app.DB.Limit(1).First(&c.Author, c.AuthorID)
	app.DB.Limit(1).First(&c.Post, c.PostID)
}

type Votes struct {
	UpVotes   uint64
	DownVotes uint64
}
