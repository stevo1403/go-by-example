package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/stevo1403/go-by-example/apps/user"
)

func LoadViews() {
	r := gin.Default()

	r.GET("/users", user.GetUsers)
	r.POST("/users", user.CreateUser)
	r.GET("/users/:id", user.GetUser)
	r.PUT("/users/:id", user.UpdateUser)
	r.DELETE("/users/:id", user.DeleteUser)

	r.Run()
}
