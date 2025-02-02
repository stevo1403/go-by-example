package post

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	user "github.com/stevo1403/go-by-example/apps/user"
	app "github.com/stevo1403/go-by-example/initializers"
)

type PostView interface {
	CreatePost()
	GetPost()
	ListPosts()
	UpdatePost()
	DeletePost()
}

func generateRandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func removeDuplicateTags(tags []string) []string {
	// Remove duplicate tags
	encountered := map[string]bool{}
	result := []string{}
	for _, tag := range tags {
		if !encountered[tag] {
			encountered[tag] = true
			result = append(result, tag)
		}
	}
	return result
}

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body PostSchema true "Post object that needs to be created"
// @Security BearerAuth
// @Success 200 {object} map[string]PostOut "{"status": "success", "data": PostOut}"
// @Router /posts [post]
func CreatePost(c *gin.Context) {

	// Parse the request body into a schema
	var postBody PostSchema
	err := c.BindJSON(&postBody)

	if err != nil {
		log.Printf("An error occurred while parsing response body: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "Invalid request body",
			"data":    map[string]interface{}{},
		})
		return
	}

	// Create a new post object from the body schema
	post := postBody.from_schema()

	// Only publish the post if it is not a draft
	if !post.IsDraft {
		post.Publish()
	}

	post.UpdateFields()

	if !post.Author.Exists((post.AuthorID)) {
		// Return an error indicating the post was not found
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "failure",
				"message": fmt.Sprintf("Author ID '%d' does not point to an existing resource.", post.AuthorID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Add tags
	postTags := removeDuplicateTags(postBody.Tags)
	// Clear the tags slice
	post.Tags = []string{}
	// Add the tags to the post
	post.Tags = append(post.Tags, postTags...)
	// Create a post object in the DB
	app.DB.Create(&post)

	// Convert comment object back to schema
	respObj := post.to_schema()

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
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
// @Success 200 {object} map[string]PostOut "{"status": "success", "data": PostOut}"
// @Failure 404 {object} map[string]interface{} "{"status": "failure", "data": {}, "message": "Post with post id '{id}' does not exist."}"
// @Router /posts/{id} [get]
func GetPost(c *gin.Context) {
	// Extract parameter 'id' from URL
	postID := c.Param("id")

	// Retrieve post from DB
	var post Post
	result := app.DB.Limit(1).First(&post, postID)
	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	fmt.Println(post)
	fmt.Print(result.Error)
	if postNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "failure",
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
				"status": "success",
				"data":   PostOut{Post: post_as_schema},
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
// @Success 200 {object} map[string]PostListOut "{"status": "success", "data": PostListOut}"
// @Router /posts [get]
func ListPosts(c *gin.Context) {
	// Retrieve posts from DB
	var posts []Post
	app.DB.Find(&posts)

	// Convert posts to schema
	posts_as_schema := []PostOutSchema{}
	for i := range posts {
		posts[i].UpdateFields()
		posts_as_schema = append(posts_as_schema, posts[i].to_schema())
	}

	// Serve post data
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "success",
			"data":   PostListOut{Posts: posts_as_schema},
		},
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
// @Success 200 {object} map[string]PostOut "{"status": "success", "data": PostOut}"
// @Failure 404 {object} map[string]interface{} "{"status": "failure", "data": {}, "message": "Post with post id '{id}' does not exist."}"
// @Router /posts/{id} [put]
func UpdatePost(c *gin.Context) {
	postID := c.Param("id")

	// Parse the request body into a schema
	var postBody PostUpdateSchema
	err := c.BindJSON(&postBody)

	if err != nil {
		log.Printf("An error occurred while parsing response body: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "Invalid request body",
			"data":    map[string]interface{}{},
		})
		return
	}

	// Query the DB for the post
	var post Post
	result := app.DB.Limit(1).First(&post, postID)

	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if postNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "failure",
				"message": fmt.Sprintf("Post with post id '%s' does not exist.", postID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Update the fields
	post.Title = postBody.Title
	post.Body = postBody.Body
	post.IsDraft = postBody.IsDraft

	// Add tags
	postTags := removeDuplicateTags(postBody.Tags)
	// Add the tags to the post
	if postTags != nil {
		post.Tags = []string{} // Clear the tags slice
		post.Tags = append(post.Tags, postTags...)
	}

	fmt.Println(post.Tags)
	// Save the updated fields
	result = app.DB.Save(&post)

	if result.Error != nil {
		log.Printf("Failed to update post: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failure",
			"message": "Failed to update post",
			"data":    map[string]interface{}{},
		})
		return
	}

	post_as_schema := post.to_schema()

	// Serve a response
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    PostOut{Post: post_as_schema},
		"message": "Post updated successfully",
	})
}

// DeletePost godoc
// @Summary Delete a post by ID
// @Description Delete a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "{"status": "success", "data": {}, "message": "Post deleted successfully"}"
// @Failure 404 {object} map[string]interface{} "{"status": "failure", "data": {}, "message": "Post with post id '{id}' does not exist."}"
// @Router /posts/{id} [delete]
func DeletePost(c *gin.Context) {
	postID := c.Param("id")

	// Query the DB for the post
	var post Post
	result := app.DB.Limit(1).First(&post, postID)

	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if postNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "failure",
				"message": fmt.Sprintf("Post with post id '%s' does not exist.", postID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Delete the post
	result = app.DB.Delete(&post)

	resultNotDeleted := (result.Error != nil)

	if resultNotDeleted {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"data":    map[string]interface{}{},
			"message": fmt.Sprintf("Post with post id '%s' could not be deleted.", postID),
		})
		return
	}

	// Serve a response
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    map[string]interface{}{},
		"message": "Post deleted successfully",
	})
}

// IncrementPostViews godoc
// @Summary Update the views of a post by ID
// @Description Increment the views of a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security BearerAuth
// @Success 200 {object} map[string]PostOut "{"status": "success", "data": PostOut, "message": "Views for post with post id '{id}' have been updated successfully."}"
// @Failure 404 {object} map[string]interface{} "{"status": "failure", "data": {}, "message": "Post with post id '{id}' does not exist."}"
// @Failure 400 {object} map[string]interface{} "{"status": "failure", "message": "something went wrong while processing your request"}"
// @Router /posts/{id}/views [put]
func IncrementPostViews(c *gin.Context) {
	postID := c.Param("id")

	// Query the DB for the post
	var post Post
	result := app.DB.Limit(1).First(&post, postID)

	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if postNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "failure",
				"message": fmt.Sprintf("Post with post id '%s' does not exist.", postID),
				"data":    map[string]interface{}{},
			},
		)
		return
	} else {
		// Update Post views
		userIDStr, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failure", "message": "something went wrong while processing your request"})
			return
		}
		userID := userIDStr.(uint)

		_user := user.User{}.GetUserByID(userID)

		// Increment the views of the post by 1
		post.IncrementViews(_user)

		// Save Post data
		app.DB.Save(&post) // Todo: Check .Error for failure

		// Serve Post data to the frontend
		post_as_schema := post.to_schema()
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  "success",
				"message": fmt.Sprintf("Views for post with post id '%s' have been updated successfully.", postID),
				"data":    PostOut{Post: post_as_schema},
			},
		)
	}
}

// UploadImage godoc
// @Summary Upload an image for a post by ID
// @Description Upload an image for a post by ID
// @Tags posts
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Post ID"
// @Param image formData file true "Image to be uploaded"
// @Param image_type formData string true "Type of image (preview or attachment)"
// @Security BearerAuth
// @Success 200 {object} PostImageOut "{"status": "success", "data": PostImageOut, "message": "Image uploaded successfully."}"
// @Failure 404 {object} map[string]interface{} "{"status": "failure", "data": {}, "message": "Post with post id '{id}' does not exist."}"
// @Failure 400 {object} map[string]interface{} "{"status": "failure", "message": "something went wrong while processing your request"}"
// @Router /posts/{id}/images [post]
func UploadImage(c *gin.Context) {
	postID := c.Param("id")

	// Query the DB for the post
	var post Post
	result := app.DB.Limit(1).First(&post, postID)

	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if postNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "failure",
				"message": fmt.Sprintf("Post with post id '%s' does not exist.", postID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Extract the image
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failure", "message": "something went wrong while processing your request"})
		return
	}

	// Get the image type from the post data
	imageType := c.PostForm("image_type")

	// Generate a unique filename
	filename, err := generateRandomHex(32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failure", "message": "failed to generate filename"})
		return
	}
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failure", "message": "Invalid post ID"})
		return
	}
	filename = fmt.Sprintf("image-%d-%s.%s", postIDInt, filename, "png")
	fileCompletePath := fmt.Sprintf("static/user-uploads/%s", filename)
	fileSavePath := fmt.Sprintf("cms/%s", fileCompletePath)

	// Save the image
	err = c.SaveUploadedFile(file, fileSavePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failure", "message": "something went wrong while processing your request"})
		return
	}

	// Ensure only one image preview and attachment is linked to the post by deleting the previews
	var postImages []PostImage
	app.DB.Where("post_id = ? AND type = ?", postID, ImageTypePreview).Find(&postImages)
	for _, image := range postImages {
		if err := app.DB.Delete(&image).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failure",
				"message": "Failed to delete existing image",
				"data":    map[string]interface{}{},
			})
			return
		}
	}

	// Create a new post image object
	postImage := PostImage{
		Type: func() ImageType {
			if imageType == "preview" {
				return ImageTypePreview
			}
			return ImageTypeAttachment
		}(),
		PostID:  post.ID,
		URL:     fileCompletePath,
		Caption: "Uploaded image",
	}
	app.DB.Save(&postImage)

	// Add the image to the post
	post.Images = append(post.Images, postImage)

	// Save the post
	app.DB.Save(&post)

	// Create response object
	postImageOut := PostImageOut{
		Image: postImage.to_schema(),
	}

	// Serve response
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "Image uploaded successfully.",
			"data":    postImageOut,
		},
	)
}

// GetImages godoc
// @Summary Get images linked to a post by ID
// @Description Get images linked to a post by ID
// @Tags posts
// @Produce json
// @Param id path string true "Post ID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "{"status": "success", "data": PostImageListOut, "message": "Images retrieved successfully."}"
// @Failure 404 {object} map[string]interface{} "{"status": "failure", "data": {}, "message": "Post with post id '{id}' does not exist."}"
// @Router /posts/{id}/images [get]
func GetImages(c *gin.Context) {
	postID := c.Param("id")

	// Query the DB for the post
	var post Post
	result := app.DB.Limit(1).First(&post, postID)

	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if postNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "failure",
				"message": fmt.Sprintf("Post with post id '%s' does not exist.", postID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Retrieve images linked to the post
	var postImages []PostImage
	app.DB.Where("post_id = ?", postID).Find(&postImages)

	// Convert images to schema
	imagesAsSchema := []PostImageOutSchema{}
	for _, image := range postImages {
		imagesAsSchema = append(imagesAsSchema, image.to_schema())
	}

	// Serve response
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "Images retrieved successfully.",
			"data":    PostImageListOut{Images: imagesAsSchema},
		},
	)
}

// GetImage godoc
// @Summary Get a specific image linked to a post by ID
// @Description Get a specific image linked to a post by ID
// @Tags posts
// @Produce json
// @Param id path string true "Post ID"
// @Param image_id path string true "Image ID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "{"status": "success", "data": PostImageOut, "message": "Image retrieved successfully."}"
// @Failure 404 {object} map[string]interface{} "{"status": "failure", "data": {}, "message": "Image with image id '{image_id}' does not exist."}"
// @Router /posts/{id}/images/{image_id} [get]
func GetImage(c *gin.Context) {
	postID := c.Param("id")
	imageID := c.Param("image_id")

	// Query the DB for the post
	var post Post
	result := app.DB.Limit(1).First(&post, postID)

	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if postNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "failure",
				"message": fmt.Sprintf("Post with post id '%s' does not exist.", postID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Query the DB for the image
	var postImage PostImage
	result = app.DB.Limit(1).First(&postImage, imageID)

	imageNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if imageNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "failure",
				"message": fmt.Sprintf("Image with image id '%s' does not exist.", imageID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Convert image to schema
	imageAsSchema := postImage.to_schema()

	// Serve response
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "Image retrieved successfully.",
			"data":    PostImageOut{Image: imageAsSchema},
		},
	)
}

// DeleteImage godoc
// @Summary Delete a specific image linked to a post by ID
// @Description Delete a specific image linked to a post by ID
// @Tags posts
// @Produce json
// @Param id path string true "Post ID"
// @Param image_id path string true "Image ID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "{"status": "success", "data": {}, "message": "Image deleted successfully."}"
// @Failure 404 {object} map[string]interface{} "{"status": "failure", "data": {}, "message": "Image with image id '{image_id}' does not exist."}"
// @Router /posts/{id}/images/{image_id} [delete]
func DeleteImage(c *gin.Context) {
	postID := c.Param("id")
	imageID := c.Param("image_id")

	// Query the DB for the post
	var post Post
	result := app.DB.Limit(1).First(&post, postID)

	postNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if postNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "failure",
				"message": fmt.Sprintf("Post with post id '%s' does not exist.", postID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Query the DB for the image
	var postImage PostImage
	result = app.DB.Limit(1).First(&postImage, imageID)

	imageNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	if imageNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "failure",
				"message": fmt.Sprintf("Image with image id '%s' does not exist.", imageID),
				"data":    map[string]interface{}{},
			},
		)
		return
	}

	// Delete the image
	result = app.DB.Delete(&postImage)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failure",
			"message": "Failed to delete image",
			"data":    map[string]interface{}{},
		})
		return
	}

	// Serve response
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "Image deleted successfully.",
			"data":    map[string]interface{}{},
		},
	)
}
