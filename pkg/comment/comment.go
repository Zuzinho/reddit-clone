package comment

import (
	"main/pkg/id"
	"main/pkg/user"
	"time"
)

type Comment struct {
	ID      string     `json:"id"`
	Body    string     `json:"body"`
	Author  *user.User `json:"author"`
	Created time.Time  `json:"created"`
}

func NewComment(body string, u *user.User) *Comment {
	return &Comment{
		ID:      id.GenerateID(),
		Body:    body,
		Author:  u,
		Created: time.Now(),
	}
}

type Comments []*Comment

func (comments *Comments) Append(comment *Comment) {
	*comments = append(*comments, comment)
}

type CommentsRepo interface {
	Create(postID string, body string, author *user.User) *Comment
	Delete(postID, commentID, userID string) error
	GetAllByPostID(postID string) *Comments
	DeleteByPostID(postID string)
}
