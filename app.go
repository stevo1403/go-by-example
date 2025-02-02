package main

import (
	"github.com/stevo1403/go-by-example/apps/comment"
	"github.com/stevo1403/go-by-example/apps/post"
	"github.com/stevo1403/go-by-example/apps/user"
	app "github.com/stevo1403/go-by-example/initializers"
)

func Migrate() {
	if err := app.DB.AutoMigrate(&user.User{}); err != nil {
		panic("Failed to migrate user model: " + err.Error())
	}
	if err := app.DB.AutoMigrate(&post.Post{}); err != nil {
		panic("Failed to migrate post model: " + err.Error())
	}
	if err := app.DB.AutoMigrate(&comment.Comment{}); err != nil {
		panic("Failed to migrate comment model: " + err.Error())
	}
	if err := app.DB.AutoMigrate(&post.PostViews{}); err != nil {
		panic("Failed to migrate post views model: " + err.Error())
	}
	if err := app.DB.AutoMigrate(&post.PostImage{}); err != nil {
		panic("Failed to migrate post image model: " + err.Error())
	}
}

func Load() {
	app.LoadEnvVariables()
	app.LoadDB()
	Migrate()
}
