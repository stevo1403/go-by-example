package views

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func LoginView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/login.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Login"}, c.Writer)
}

func SignUpView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/signup.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Sign Up"}, c.Writer)
}

func HomeView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/home.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Home"}, c.Writer)
}

func ProfileView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/profile.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Profile"}, c.Writer)
}

func PostListView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/posts.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Posts"}, c.Writer)
}

func CommentListView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/comments.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Comments"}, c.Writer)
}

func MediaView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/media.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Media"}, c.Writer)
}

func SettingsView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/settings.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Settings"}, c.Writer)
}
