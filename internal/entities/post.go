package entities

type Post struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var PostsFakeDb = &[]Post{
	{ID: 1, Title: "First Post", Content: "This is the content of the first post."},
	{ID: 2, Title: "Second Post", Content: "This is the content of the second post."},
}
