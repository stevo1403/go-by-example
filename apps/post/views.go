package post

import (
	"fmt"
	"log"
	"net/http"

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
	post.UpdateFields()

	if !post.Author.Exists((post.AuthorID)) {
		// Return an error indicating the post was not found
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("Author ID '%d' does not point to an existing resource.", post.AuthorID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Create a post object in the DB
	app.DB.Create(&post)

	// Convert comment object back to schema
	respObj := post.to_schema()

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Post created successfully",
			"data":    PostOut{Post: respObj},
		},
	)

}
func GetPost(c *gin.Context) {
	// Extract parameter 'id' from URL
	postID := c.Param("id")

	// Retrieve post from DB
	var post Post
	result := app.DB.Limit(1).First(&post, postID)
	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if postNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("Post with post id '%s' does not exist.", postID),
				"data":    map[string]interface{}{},
			},
		)
		return
	} else {

		post_as_schema := post.to_schema()
		c.JSON(
			http.StatusOK,
			gin.H{
				"data": PostOut{Post: post_as_schema},
			},
		)
	}

}
func ListPosts(c *gin.Context) {
	// Retrieve posts from DB
	var posts []Post
	app.DB.Find(&posts)

	// Convert posts to schema
	var posts_as_schema []PostOutSchema
	for _, post := range posts {
		posts_as_schema = append(posts_as_schema, post.to_schema())
	}

	// Serve post data
	c.JSON(
		http.StatusOK,
		gin.H{"data": PostListOut{Posts: posts_as_schema}},
	)
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
	result := app.DB.Limit(1).First(&post, postID)

	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if postNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("Post with post id '%s' does not exist.", postID),
				"data":    map[string]interface{}{},
			},
		)
		return
	} else {
		// Update Post data
		post.Title = postBody.Title
		post.Body = postBody.Body

		// Save Post data
		app.DB.Save(&post) // Todo: Check .Error for failure

		// Serve Post data to the frontend
		post_as_schema := post.to_schema()
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": fmt.Sprintf("Post with post id '%s' has been updated successfully.", postID),
				"data":    PostOut{Post: post_as_schema},
			},
		)
	}

}
func DeletePost(c *gin.Context) {
	// Extract parameter `id` from URL
	postID := c.Param("id")

	// Query DB for post object
	var post Post
	result := app.DB.Limit(1).First(&post, postID)

	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)
	if postNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("Post with post id '%s' does not exist.", postID),
				"data":    map[string]interface{}{},
			},
		)
		return
	} else {
		// Delete post from DB
		result = app.DB.Delete(&post)
		resultNotDeleted := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

		if resultNotDeleted {
			c.JSON(http.StatusNotFound, gin.H{
				"data":    map[string]interface{}{},
				"message": fmt.Sprintf("Post with post id '%s' could not be deleted.", postID),
			})
			return
		}
		// Serve a response
		c.JSON(http.StatusOK, gin.H{
			"data":    map[string]interface{}{},
			"message": "Post deleted successfully",
		})
		return
	}

}
