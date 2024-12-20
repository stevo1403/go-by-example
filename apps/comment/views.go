package comment

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	app "github.com/stevo1403/go-by-example/initializers"
)

// Create views for comments

type CommentView interface {
	CreateComment()
	ListComments()
	GetComment()
	UpdateComment()
	DeleteComment()
	UpVoteComment()
	DownVoteComment()
}

// CreateComment godoc
// @Summary Create a new comment
// @Description Create a new comment
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body CommentSchema true "Comment object that needs to be created"
// @Security BearerAuth
// @Success 200 {object} map[string]CommentOut "{"data": CommentOut}"
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	var commentBody CommentSchema

	// Store the request body in `commentBody``
	err := c.BindJSON(&commentBody)

	if err != nil {
		log.Fatalf("An error occurred while converting parsing response body: %s", err)
	}

	// Create a new comment object from the body schema
	comment := commentBody.from_schema()
	// comment.UpdateFields()
	comment.UpdateFields()

	// Check if AuthorID points to a real author(user)
	if !comment.Author.Exists(comment.AuthorID) {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": fmt.Sprintf("Author ID '%d' does not point to an existing resource.", comment.AuthorID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Check if PostID points to a real post
	if !comment.Post.Exists(comment.PostID) {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": fmt.Sprintf("Post ID '%d' does not point to an existing resource.", comment.PostID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Create a comment object in the database
	app.DB.Create(&comment)

	// Convert comment object back to schema
	respObj := comment.to_schema()

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Comment created successfully",
			"data":    CommentOut{Comment: respObj},
		},
	)
}

// ListComments godoc
// @Summary List all comments
// @Description List all comments
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]CommentListOut "{"data": CommentListOut}"
// @Router /comments [get]
func ListComments(c *gin.Context) {
	var comments []Comment

	// Get a list of `Comment` objects from the database
	app.DB.Find(&comments)

	// Convert `Comment` objects to schemas
	var comments_schema []CommentOutSchema
	for _, comment := range comments {
		comments_schema = append(comments_schema, comment.to_schema())
	}

	// Serve the converted data as response to the frontend
	c.JSON(
		http.StatusOK,
		gin.H{
			"data": CommentListOut{Comments: comments_schema},
		},
	)
}

// GetComment godoc
// @Summary Get a comment by ID
// @Description Get a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Security BearerAuth
// @Success 200 {object} map[string]CommentOut "{"data": CommentOut}"
// @Failure 404 {object} map[string]interface{} "{"data": {}, "message": "Comment with comment id '{id}' does not exist."}"
// @Router /comments/{id} [get]
func GetComment(c *gin.Context) {
	// Extract parameter `id` from URL
	commentID := c.Param("id")

	// Query the DB for comment with id `commentID`
	var comment Comment
	result := app.DB.Limit(1).First(&comment, commentID)

	commentNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if commentNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("The comment identified by comment id '%s' could not be found.", commentID),
				"data":    map[string]interface{}{},
			},
		)
	} else {

		// Convert comment to schema
		comment_as_schema := comment.to_schema()

		// Serve schema as response
		c.JSON(
			http.StatusOK,
			gin.H{
				"data": CommentOut{Comment: comment_as_schema},
			},
		)
	}
}

// UpdateComment godoc
// @Summary Update a comment by ID
// @Description Update a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Param comment body CommentUpdateSchema true "Comment object that needs to be updated"
// @Security BearerAuth
// @Success 200 {object} map[string]CommentOut "{"data": CommentOut}"
// @Failure 404 {object} map[string]interface{} "{"data": {}, "message": "Comment with comment id '{id}' does not exist."}"
// @Router /comments/{id} [put]
func UpdateComment(c *gin.Context) {
	// Extract parameter `id` from URL
	commentID := c.Param("id")

	// Convert the request body to a schema
	var commentBody CommentUpdateSchema
	err := c.BindJSON(&commentBody)

	if err != nil {
		log.Fatalf("An error occurred while converting parsing response body: %s", err)
	}

	// Find the comment object with the given comment id
	var comment Comment
	result := app.DB.Limit(1).First(&comment, commentID)

	commentNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if commentNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("The comment identified by comment id '%s' could not be found.", commentID),
				"data":    map[string]interface{}{},
			},
		)
	} else {

		// Update comment object
		comment.Body = commentBody.Body

		// Save the updated data
		app.DB.Save(&comment)

		// Convert the comment object to a schema
		comment_as_schema := comment.to_schema()

		c.JSON(
			http.StatusOK,
			gin.H{
				"message": fmt.Sprintf("The comment with comment id '%s' was updated successfully.", commentID),
				"data":    CommentOut{Comment: comment_as_schema},
			},
		)
	}

}

// DeleteComment godoc
// @Summary Delete a comment by ID
// @Description Delete a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "{"data": {}, "message": "Comment deleted successfully"}"
// @Failure 404 {object} map[string]interface{} "{"data": {}, "message": "Comment with comment id '{id}' does not exist."}"
// @Router /comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	// Extract parameter `id` from URL
	commentID := c.Param("id")

	// Find the comment object with the given comment id
	var comment Comment
	result := app.DB.Limit(1).First(&comment, commentID)

	commentNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if commentNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' could not be found.", commentID),
				"data":    map[string]interface{}{},
			},
		)
	} else {
		// Delete comment from DB
		result = app.DB.Delete(&comment)
		resultNotDeleted := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

		if resultNotDeleted {
			c.JSON(http.StatusNotFound, gin.H{
				"data":    map[string]interface{}{},
				"message": fmt.Sprintf("An error occurred: comment identified by comment id '%s' could not be deleted.", commentID),
			})
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' deleted successfully.", commentID),
				"data":    map[string]interface{}{},
			},
		)
	}

}

// UpvoteComment godoc
// @Summary Upvote a comment by ID
// @Description Upvote a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Security BearerAuth
// @Success 200 {object} map[string]CommentOut "{"data": CommentOut}"
// @Failure 404 {object} map[string]interface{} "{"data": {}, "message": "Comment with comment id '{id}' does not exist."}"
// @Router /comments/{id}/upvote [put]
func UpvoteComment(c *gin.Context) {
	// Extract parameter `id` from URL
	commentID := c.Param("id")

	// Query the DB for comment with id `commentID`
	var comment Comment
	result := app.DB.Limit(1).First(&comment, commentID)

	commentNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if commentNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' could not be found.", commentID),
				"data":    map[string]interface{}{},
			},
		)
	} else {
		// Increase the upvotes
		comment.UpVotes = comment.UpVotes + 1

		// Save the updated field
		app.DB.Save(&comment)

		// Convert comment to schema
		comment_as_schema := comment.to_schema()

		c.JSON(
			http.StatusOK,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' has been successfully upvoted.", commentID),
				"data":    CommentOut{Comment: comment_as_schema},
			},
		)
	}

}

// DownVoteComment godoc
// @Summary Downvote a comment by ID
// @Description Downvote a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Security BearerAuth
// @Success 200 {object} map[string]CommentOut "{"data": CommentOut}"
// @Failure 404 {object} map[string]interface{} "{"data": {}, "message": "Comment with comment id '{id}' does not exist."}"
// @Router /comments/{id}/downvote [put]
func DownVoteComment(c *gin.Context) {
	// Extract parameter `id` from URL
	commentID := c.Param("id")

	// Query the DB for comment with id `commentID`
	var comment Comment
	result := app.DB.Limit(1).First(&comment, commentID)

	commentNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if commentNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' could not be found.", commentID),
				"data":    map[string]interface{}{},
			},
		)
	} else {
		// Increase the downvotes or descrease the upvotes
		comment.DownVotes++

		// Save the updated field
		app.DB.Save(&comment)

		// Convert comment to schema
		comment_as_schema := comment.to_schema()

		c.JSON(
			http.StatusOK,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' has been successfully upvoted.", commentID),
				"data":    CommentOut{Comment: comment_as_schema},
			},
		)
	}

}
