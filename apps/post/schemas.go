package post

type PostSchema struct {
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	AuthorID uint     `json:"author_id"`
	Status   string   `json:"status"`
	Tags     []string `json:"tags"`
}

type PostOutSchema struct {
	ID          uint     `json:"id"`
	AuthorID    uint     `json:"author_id"`
	AuthorName  string   `json:"author_name"`
	Title       string   `json:"title"`
	Body        string   `json:"body"`
	Views       uint64   `json:"views"`
	PublishedAt string   `json:"published_at"`
	IsDraft     bool     `json:"is_draft"`
	Tags        []string `json:"tags"`
}

type PostOut struct {
	Post PostOutSchema `json:"post"`
}

type PostUpdateSchema struct {
	Title   string   `json:"title"`
	Body    string   `json:"body"`
	IsDraft bool     `json:"is_draft"`
	Tags    []string `json:"tags"`
}

func (p *Post) to_schema() PostOutSchema {
	return PostOutSchema{
		ID:          p.ID,
		AuthorID:    p.AuthorID,
		AuthorName:  p.Author.GetFullName(),
		Title:       p.Title,
		Body:        p.Body,
		Views:       p.Views,
		PublishedAt: p.PublishedAt.Format("2006-01-02 15:04:05"),
		IsDraft:     p.IsDraft,
		Tags: func() []string {
			if p.Tags == nil {
				return []string{}
			} else {
				tags := make([]string, len(p.Tags))
				for i, tag := range p.Tags {
					tags[i] = string(tag)
				}
				return tags
			}
		}(),
	}
}

func (_s PostSchema) from_schema() Post {
	return Post{
		Title:    _s.Title,
		Body:     _s.Body,
		AuthorID: _s.AuthorID,
		IsDraft:  _s.Status == "draft",
		Tags:     _s.Tags,
	}
}

type PostListOut struct {
	Posts []PostOutSchema `json:"posts"`
}

type PostImageSchema struct {
	ID      uint   `json:"id"`
	PostID  uint   `json:"post_id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Caption string `json:"caption"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

type PostImageIn struct {
	Image string `json:"image"`
}

type PostImageOutSchema struct {
	ID     uint   `json:"id"`
	PostID uint   `json:"post_id"`
	URL    string `json:"url"`
	Type   string `json:"image_type"`
}

func (pi *PostImage) to_schema() PostImageOutSchema {
	return PostImageOutSchema{
		ID:     pi.ID,
		PostID: pi.PostID,
		URL:    pi.URL,
		Type:   string(pi.Type),
	}
}

type PostImageOut struct {
	Image PostImageOutSchema `json:"image"`
}

type PostImageListOut struct {
	Images []PostImageOutSchema `json:"images"`
}
