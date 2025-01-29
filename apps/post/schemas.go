package post

type PostSchema struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorID uint   `json:"author_id"`
}

type PostOutSchema struct {
	ID          uint   `json:"id"`
	AuthorID    uint   `json:"author_id"`
	AuthorName  string `json:"author_name"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Views       uint64 `json:"views"`
	PublishedAt string `json:"published_at"`
}

type PostOut struct {
	Post PostOutSchema `json:"post"`
}

type PostUpdateSchema struct {
	Title string `json:"title"`
	Body  string `json:"body"`
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
	}
}

func (_s PostSchema) from_schema() Post {
	return Post{
		Title:    _s.Title,
		Body:     _s.Body,
		AuthorID: _s.AuthorID,
	}
}

type PostListOut struct {
	Posts []PostOutSchema `json:"posts"`
}
