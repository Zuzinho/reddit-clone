package comment

import (
	"main/pkg/id"
	"main/pkg/user"
	"time"
)

// Comment - тип комментария
type Comment struct {
	ID      string     `json:"id"`
	Body    string     `json:"body"`
	Author  *user.User `json:"author"`
	Created time.Time  `json:"created"`
}

// NewComment создает тип Comment
func NewComment(body string, u *user.User) *Comment {
	return &Comment{
		ID:      id.GenerateID(),
		Body:    body,
		Author:  u,
		Created: time.Now(),
	}
}

// Comments - массив ссылок на Comment
type Comments []*Comment

// Append добавляет в Comments ссылку на Comment
func (comments *Comments) Append(comment *Comment) {
	*comments = append(*comments, comment)
}

// CommentsRepo - интерфейс для хранения Comment
type CommentsRepo interface {
	Create(postID string, body string, author *user.User) *Comment
	Delete(postID, commentID, userID string) error
	GetAllByPostID(postID string) *Comments
	DeleteByPostID(postID string)
}
