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

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body PostSchema true "Post object that needs to be created"
// @Security BearerAuth
// @Success 200 {object} map[string]PostOut "{"data": PostOut}"
// @Router /posts [post]
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

// GetPost godoc
// @Summary Get a post by ID
// @Description Get a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security BearerAuth
// @Success 200 {object} map[string]PostOut "{"data": PostOut}"
// @Failure 404 {object} map[string]interface{} "{"data": {}, "message": "Post with post id '{id}' does not exist."}"
// @Router /posts/{id} [get]
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

// ListPosts godoc
// @Summary Get all posts
// @Description Get all posts
// @Tags posts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]PostListOut "{"data": PostListOut}"
// @Router /posts [get]
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

// UpdatePost godoc
// @Summary Update a post by ID
// @Description Update a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param post body PostUpdateSchema true "Post object that needs to be updated"
// @Security BearerAuth
// @Success 200 {object} map[string]PostOut "{"data": PostOut}"
// @Failure 404 {object} map[string]interface{} "{"data": {}, "message": "Post with post id '{id}' does not exist."}"
// @Router /posts/{id} [put]
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
// DeletePost godoc
// @Summary Delete a post by ID
// @Description Delete a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "{"data": {}, "message": "Post deleted successfully"}"
// @Failure 404 {object} map[string]interface{} "{"data": {}, "message": "Post with post id '{id}' does not exist."}"
// @Router /posts/{id} [delete]
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
