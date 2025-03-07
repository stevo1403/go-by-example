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

func ResetPasswordView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/password-reset.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Reset Password"}, c.Writer)
}

func GetPostView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/post.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Blog Post"}, c.Writer)
}

func EditPostView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/edit-post.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Edit Blog Post"}, c.Writer)
}

func EditCommentView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/edit-comment.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Edit Comment"}, c.Writer)
}

func GetCommentView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/comment.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Comment to Blog Post"}, c.Writer)
}

func CreatePostView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/create-post.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Create Post"}, c.Writer)
}

func CreateCommentView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/create-comment.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Create Comment"}, c.Writer)
}

func CreateMediaView(c *gin.Context) {
	tmpl := pongo2.Must(pongo2.FromFile("cms/templates/create-media.html"))
	c.Writer.WriteHeader(http.StatusOK)
	tmpl.ExecuteWriter(pongo2.Context{"title": "Create Media"}, c.Writer)
}
