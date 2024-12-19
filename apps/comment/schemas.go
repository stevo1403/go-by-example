package comment

type CommentSchema struct {
	AuthorID  uint   `json:"author_id"`
	PostID    uint   `json:"post_id"`
	Body      string `json:"body"`
	UpVotes   uint64 `json:"upvotes"`
	DownVotes uint64 `json:"downvotes"`
}

type CommentOutSchema struct {
	ID         uint   `json:"id"`
	AuthorID   uint   `json:"author_id"`
	AuthorName string `json:"author_name"`
	PostID     uint   `json:"post_id"`
	Body       string `json:"body"`
	UpVotes    uint64 `json:"upvotes"`
	DownVotes  uint64 `json:"downvotes"`
}

type CommentOut struct {
	Comment CommentOutSchema `json:"comment"`
}

// Converts a `Comment“ object to a `CommentOutSchema“ object
func (c Comment) to_schema() CommentOutSchema {
	return CommentOutSchema{
		ID:         c.ID,
		AuthorID:   c.AuthorID,
		AuthorName: c.Author.GetUserByID(c.AuthorID).GetFullName(),
		PostID:     c.PostID,
		Body:       c.Body,
		UpVotes:    c.UpVotes,
		DownVotes:  c.DownVotes,
	}
}

// Convert `CommentSchema` to `Comment` object
func (_s *CommentSchema) from_schema() Comment {
	return Comment{
		AuthorID: _s.AuthorID,
		PostID:   _s.PostID,
		Body:     _s.Body,
	}
}

type CommentListOut struct {
	// List of comment objects
	Comments []CommentOutSchema `json:"comments"`
}

type CommentUpdateSchema struct {
	Body string `json:"body"`
}
