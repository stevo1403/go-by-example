package main

import (
	"github.com/stevo1403/go-by-example/apps/comment"
	"github.com/stevo1403/go-by-example/apps/post"
	"github.com/stevo1403/go-by-example/apps/user"
	app "github.com/stevo1403/go-by-example/initializers"
)

func Migrate() {
	app.DB.AutoMigrate(&user.User{})
	app.DB.AutoMigrate(&post.Post{})
	app.DB.AutoMigrate(&comment.Comment{})
}

func Load() {
	app.LoadEnvVariables()
	app.LoadDB()
	Migrate()
}
