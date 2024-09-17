package post

import (
	"github.com/stevo1403/go-by-example/apps/user"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string
	Body     string
	AuthorID int       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign key
	Author   user.User // Todo: Change this to allow multiple author per posts
}
