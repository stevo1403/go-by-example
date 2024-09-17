package post

type PostSchema struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorID int    `json:"author_id"`
}

type PostOutSchema struct {
	ID         uint   `json:"id"`
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	Title      string `json:"title"`
	Body       string `json:"body"`
}

type PostOut struct {
	Post PostOutSchema `json:"post"`
}

type PostUpdateSchema struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (p Post) to_schema() PostOutSchema {
	return PostOutSchema{
		ID:         p.ID,
		AuthorID:   p.AuthorID,
		AuthorName: p.Author.GetFullName(),
		Title:      p.Title,
		Body:       p.Body,
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
	Posts []PostOutSchema
}

// type PostListOut type SortBy []Type

// func (a SortBy) Len() int           { return len(a) }
// func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }{

// }
