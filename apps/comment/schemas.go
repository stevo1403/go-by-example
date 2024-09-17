package comment

type CommentSchema struct {
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	PostID     int    `json:"post_id"`
	Body       string `json:"body"`
	UpVotes    int    `json:"upvotes"`
	DownVotes  int    `json:"downvotes"`
}

type CommentOutSchema struct {
	ID         uint   `json:"id"`
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	PostID     int    `json:"post_id"`
	Body       string `json:"body"`
	UpVotes    int    `json:"upvotes"`
	DownVotes  int    `json:"downvotes"`
}

type CommentOut struct {
	Comment CommentOutSchema `json:"comment"`
}

// Converts a `Comment“ object to a `CommentOutSchema“ object
func (c Comment) to_schema() CommentOutSchema {
	return CommentOutSchema{
		ID:         c.ID,
		AuthorID:   c.AuthorID,
		AuthorName: c.Author.FirstName + " " + c.Author.LastName,
		PostID:     c.PostID,
		Body:       c.Body,
		UpVotes:    c.UpVotes,
		DownVotes:  c.DownVotes,
	}
}

// Convert `CommentSchema` to `Comment` object
func (_s *CommentSchema) from_schema() Comment {
	return Comment{
		AuthorID:  _s.AuthorID,
		PostID:    _s.PostID,
		Body:      _s.Body,
		UpVotes:   _s.UpVotes,
		DownVotes: _s.DownVotes,
	}
}

type CommentListOut struct {
	// List of comment objects
	Comments []CommentOutSchema `json:"comments"`
}

type CommentUpdateSchema struct {
	Body string `json:"body"`
}
