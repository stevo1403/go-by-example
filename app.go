package main

import (
	"github.com/stevo1403/go-by-example/apps/comment"
	"github.com/stevo1403/go-by-example/apps/post"
	"github.com/stevo1403/go-by-example/apps/user"
	"github.com/stevo1403/go-by-example/initializers"
)

func Migrate() {
	initializers.DB.AutoMigrate(&user.User{})
	initializers.DB.AutoMigrate(&post.Post{})
	initializers.DB.AutoMigrate(&comment.Comment{})
}

func Load() {
	initializers.LoadEnvVariables()
	initializers.LoadDB()
	Migrate()
}
