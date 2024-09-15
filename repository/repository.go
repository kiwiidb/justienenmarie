package repository

type Repository struct {
}

func (r *Repository) CreatePost(post Post) (string, error) {
	return "", nil
}

func (r *Repository) GetPost(id string) (Post, error) {
	return Post{}, nil
}

func (r *Repository) GetPosts() ([]Post, error) {
	return []Post{}, nil
}

func (r *Repository) CreateAnswer(answer Answer) (string, error) {
	return "", nil
}

func (r *Repository) UpdateAnswer(id string) (Answer, error) {
	return Answer{}, nil
}

func (r *Repository) GetAnswers() (Answer, error) {
	return Answer{}, nil
}
