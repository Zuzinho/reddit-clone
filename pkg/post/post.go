package post

import (
	"main/pkg/id"
	"main/pkg/user"
	"time"
)

// Post - тип поста
type Post struct {
	ID       string     `json:"id"`
	Author   *user.User `json:"author"`
	Title    string     `json:"title"`
	Text     string     `json:"text"`
	Category Category   `json:"category"`
	Type     Type       `json:"type"`
	Created  time.Time  `json:"created"`
	Views    uint32     `json:"views"`
}

// NewPost создает экземпляр Post
func NewPost(author *user.User, title, text string, category Category, tp Type) (*Post, error) {
	err := category.Valid()
	if err != nil {
		return nil, err
	}

	err = tp.Valid()
	if err != nil {
		return nil, err
	}

	return &Post{
		ID:       id.GenerateID(),
		Author:   author,
		Title:    title,
		Text:     text,
		Category: category,
		Type:     tp,
		Created:  time.Now(),
		Views:    0,
	}, nil
}

// Posts - массив ссылок на Post
type Posts []*Post

// Append добавляет Post в Posts
func (posts *Posts) Append(post *Post) {
	*posts = append(*posts, post)
}

// PostsRepo - интерфейс для хранения Post
type PostsRepo interface {
	Create(author *user.User, title, text string, category Category, tp Type) (*Post, error)
	Get(postID string) (*Post, error)
	GetAll() *Posts
	GetAllByUserName(userName string) *Posts
	GetAllByCategory(category Category) (*Posts, error)
	Delete(postID, userID string) error
}
