package comment

import (
	"fmt"
	"log"

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

func CreateComment(c *gin.Context) {
	var commentBody CommentSchema

	// Store the request body in `commentBody``
	err := c.BindJSON(&commentBody)

	if err != nil {
		log.Fatalf("An error occurred while converting parsing response body: %s", err)
	}

	// Create a new comment object from the body schema
	comment := commentBody.from_schema()

	// Create a comment object in the database
	app.DB.Create(&comment)

	// Convert comment object back to schema
	respObj := comment.to_schema()

	c.JSON(
		200,
		gin.H{
			"message": "Comment created successfully",
			"data":    CommentOut{Comment: respObj},
		},
	)
}

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
		200,
		gin.H{
			"data": CommentListOut{Comments: comments_schema},
		},
	)
}

func GetComment(c *gin.Context) {
	// Extract parameter `id` from URL
	commentID := c.Param("id")

	// Query the DB for comment with id `commentID`
	var comment Comment
	result := app.DB.Find(&comment, commentID)

	commentNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if commentNotFound {
		c.JSON(
			200,
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
			200,
			gin.H{
				"data": CommentOut{Comment: comment_as_schema},
			},
		)
	}
}

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
	result := app.DB.Find(&comment, commentID)

	commentNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if commentNotFound {
		c.JSON(
			200,
			gin.H{
				"message": fmt.Sprintf("The comment identified by comment id '%s' could not be found.", commentID),
				"data":    map[string]interface{}{},
			},
		)
	} else {

		// Update comment object
		comment.Body = commentBody.Body

		// Save the updated data
		app.DB.Save(comment)

		// Convert the comment object to a schema
		comment_as_schema := comment.to_schema()

		c.JSON(
			200,
			gin.H{
				"message": fmt.Sprintf("The comment with comment id '%s' was updated successfully.", commentID),
				"data":    CommentOut{Comment: comment_as_schema},
			},
		)
	}

}
func DeleteComment(c *gin.Context) {
	// Extract parameter `id` from URL
	commentID := c.Param("id")

	// Find the comment object with the given comment id
	var comment Comment
	result := app.DB.Find(&comment, commentID)

	commentNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if commentNotFound {
		c.JSON(
			200,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' could not be found.", commentID),
				"data":    map[string]interface{}{},
			},
		)
	} else {
		c.JSON(
			200,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' deleted successfully.", commentID),
				"data":    map[string]interface{}{},
			},
		)
	}

}

func UpvoteComment(c *gin.Context) {
	// Extract parameter `id` from URL
	commentID := c.Param("id")

	// Query the DB for comment with id `commentID`
	var comment Comment
	result := app.DB.Find(&comment, commentID)

	commentNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if commentNotFound {
		c.JSON(
			200,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' could not be found.", commentID),
				"data":    map[string]interface{}{},
			},
		)
	} else {
		// Increase the upvotes
		comment.UpVotes = comment.UpVotes + 1

		// Save the updated field
		app.DB.Save(comment)

		// Convert comment to schema
		comment_as_schema := comment.to_schema()

		c.JSON(
			200,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' has been successfully upvoted.", commentID),
				"data":    CommentOut{Comment: comment_as_schema},
			},
		)
	}

}

func DownVoteComment(c *gin.Context) {
	// Extract parameter `id` from URL
	commentID := c.Param("id")

	// Query the DB for comment with id `commentID`
	var comment Comment
	result := app.DB.Find(&comment, commentID)

	commentNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if commentNotFound {
		c.JSON(
			200,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' could not be found.", commentID),
				"data":    map[string]interface{}{},
			},
		)
	} else {
		// Increase the downvotes or descrease the upvotes
		comment.DownVotes++

		// Save the updated field
		app.DB.Save(comment)

		// Convert comment to schema
		comment_as_schema := comment.to_schema()

		c.JSON(
			200,
			gin.H{
				"message": fmt.Sprintf("Comment identified by comment id '%s' has been successfully upvoted.", commentID),
				"data":    CommentOut{Comment: comment_as_schema},
			},
		)
	}

}
