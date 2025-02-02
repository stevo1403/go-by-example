package post

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"

	"github.com/stevo1403/go-by-example/apps/user"
	app "github.com/stevo1403/go-by-example/initializers"
)

type Post struct {
	gorm.Model
	Title       string      // Title of the post
	Body        string      // Body of the post
	Views       uint64      // Number of views
	IsDraft     bool        `gorm:"default:false"` // Is the post a draft, should be false by default
	PublishedAt time.Time   // Date the post was published
	AuthorID    uint        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign key
	Author      user.User   // Todo: Change this to allow multiple author per posts
	mu          sync.Mutex  // Mutex for incrementing views
	Tags        PostTag     `json:"tags" gorm:"type:text"`
	Images      []PostImage `json:"images" gorm:"foreignKey:PostID"`
}

type PostTag []string

func (s *PostTag) Scan(value interface{}) error {
	if value == nil {
		*s = nil
		return nil
	}

	var bytes []byte

	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}

	return json.Unmarshal(bytes, s)
}

func (s PostTag) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type PostImage struct {
	gorm.Model
	Type    ImageType `json:"type"` // Type of the image (preview, attachment, etc)
	PostID  uint      `json:"post_id"`
	URL     string    `json:"url"`     // Path to the uploaded image
	Caption string    `json:"caption"` // Optional caption for the image
	Width   int       `json:"width"`   // Optional image width
	Height  int       `json:"height"`  // Optional image height
}

type PostViews struct {
	gorm.Model
	PostID uint // Foreign key for Post
	Post
	Viewers []user.User `gorm:"many2many:post_view_users;"` // Many-to-many relationship with User
}

type ImageType string

const (
	ImageTypePreview    ImageType = "preview"
	ImageTypeAttachment ImageType = "attachment"
)

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

	p.Views++                                                           // Increment the views
	app.DB.FirstOrCreate(&postViews, PostViews{Post: *p, PostID: p.ID}) // Create a new record if it does not exist
	postViews.Viewers = append(postViews.Viewers, Viewer)               // Append the viewer to the viewers list
	app.DB.Save(&postViews)                                             // Save the post views
}
