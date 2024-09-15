package repository

// a post has a a videolink, a question, and a string id

type Post struct {
	ID        string `json:"id"`
	Videolink string `json:"videolink"`
	Question  string `json:"question"`
}

type Answer struct {
	ID       string `json:"id"`
	PostID   string `json:"postId"`
	Username string `json:"username"`
	Text     string `json:"text"`
	Correct  *bool  `json:"correct"`
}
