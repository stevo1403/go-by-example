package post

import (
	"gorm.io/gorm"

	"github.com/stevo1403/go-by-example/apps/user"
	app "github.com/stevo1403/go-by-example/initializers"
)

type Post struct {
	gorm.Model
	Title    string
	Body     string
	AuthorID uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign key
	Author   user.User // Todo: Change this to allow multiple author per posts
}

func (p Post) UpdateFields() {
	app.DB.Limit(1).First(&p.Author, p.AuthorID)
}

func (p Post) Exists(postID uint) bool {
	// Checks if post exists
	var post Post
	result := app.DB.Limit(1).First(&post, postID)
	recordNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	return !recordNotFound

}
