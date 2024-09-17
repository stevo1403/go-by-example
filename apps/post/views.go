package post

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	app "github.com/stevo1403/go-by-example/initializers"
)

type PostView interface {
	CreatePost()
	GetPost()
	ListPost()
	UpdatePost()
	DeletePost()
}

func CreatePost(c *gin.Context) {

	// Parse the request body into a schema
	var postBody PostSchema
	err := c.BindJSON(&postBody)

	if err != nil {
		log.Fatalf("An error occurred while converting parsing response body: %s", err)
	}

	// Create a new post object from the body schema
	post := postBody.from_schema()

	// Create a post object in the DB
	app.DB.Create(&post)

	// Convert comment object back to schema
	respObj := post.to_schema()

	c.JSON(
		200,
		gin.H{
			"message": "Post created successfully",
			"data":    PostOut{Post: respObj},
		},
	)

}

func GetPost(c *gin.Context) {

}
func ListPost(c *gin.Context) {

}
func UpdatePost(c *gin.Context) {

	postID := c.Param("id")

	// Parse the request body into a schema
	var postBody PostUpdateSchema
	err := c.BindJSON(&postBody)

	if err != nil {
		log.Fatalf("An error occurred while converting parsing response body: %s", err)
	}

	// Query the DB for the post
	var post Post
	result := app.DB.Find(&post, postID)

	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if postNotFound {
		c.JSON(
			200,
			gin.H{
				"message": fmt.Sprintf("Post with post id '%s' does not exist.", postID),
				"data":    map[string]interface{}{},
			},
		)
	} else {
		// Update Post data
		post.Title = postBody.Title
		post.Body = postBody.Body

		// Save Post data
		app.DB.Save(post) // Todo: Check .Error for failure

		// Serve Post data to the frontend
		post_as_schema := post.to_schema()
		c.JSON(
			200,
			gin.H{
				"message": fmt.Sprintf("Post with post id '%s' has been updated successfully.", postID),
				"data":    PostOut{Post: post_as_schema},
			},
		)
	}

}
func DeletePost(c *gin.Context) {

}
