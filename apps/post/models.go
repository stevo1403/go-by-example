package post

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"

	"github.com/stevo1403/go-by-example/apps/user"
	app "github.com/stevo1403/go-by-example/initializers"
)

type Post struct {
	gorm.Model
	Title       string     // Title of the post
	Body        string     // Body of the post
	Views       uint64     // Number of views
	IsDraft     bool       `gorm:"default:false"` // Is the post a draft, should be false by default
	PublishedAt time.Time  // Date the post was published
	AuthorID    uint       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign key
	Author      user.User  // Todo: Change this to allow multiple author per posts
	mu          sync.Mutex // Mutex for incrementing views
}

type PostViews struct {
	gorm.Model
	PostID uint // Foreign key for Post
	Post
	Viewers []user.User `gorm:"many2many:post_view_users;"` // Many-to-many relationship with User
}

func (p *Post) UpdateFields() {
	app.DB.Limit(1).First(&p.Author, p.AuthorID)
}

func (p *Post) Exists(postID uint) bool {
	// Checks if post exists
	var post Post
	result := app.DB.Limit(1).First(&post, postID)
	recordNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	return !recordNotFound

}

// Publish the post
func (p *Post) Publish() {
	if p.IsDraft {
		p.IsDraft = false
	} else {
		p.PublishedAt = time.Now()
	}
}

// IncrementViews safely increments the Views field
func (p *Post) IncrementViews(Viewer user.User) {
	var postViews PostViews
	isViewed := app.DB.Limit(1).
		Where(map[string]interface{}{"post_id": p.ID, "author_id": Viewer.ID}).
		First(&postViews).
		Error == nil
	if isViewed {
		return
	}

	p.mu.Lock() // Lock the mutex
	defer p.mu.Unlock()

	p.Views++                                                                   // Increment the views
	fmt.Println("Incrementing views for post with ID: ", p.ID, " to ", p.Views) // Log the increment
	app.DB.FirstOrCreate(&postViews, PostViews{Post: *p, PostID: p.ID})         // Create a new record if it does not exist
	postViews.Viewers = append(postViews.Viewers, Viewer)                       // Append the viewer to the viewers list
	app.DB.Save(&postViews)                                                     // Save the post views
}
