package comment

import "fmt"

// NoCommentError - Ошибка для отсутствия Comment
type NoCommentError struct {
	commentID string
}

func newNoCommentError(commentID string) NoCommentError {
	return NoCommentError{
		commentID: commentID,
	}
}

// Error возвращает текстовое представление ошибки
func (err NoCommentError) Error() string {
	return fmt.Sprintf("no comment '%s'", err.commentID)
}

// NoRightsError - ошибка отсутствия прав
type NoRightsError struct {
	commentID string
}

func newNoRightsError(commentID string) NoRightsError {
	return NoRightsError{
		commentID: commentID,
	}
}

// Error возвращает текстовое представление ошибки
func (err NoRightsError) Error() string {
	return fmt.Sprintf("have no rights for comment '%s'", err.commentID)
}

var (
	NoCommentErr = NoCommentError{}
	NoRightsErr  = NoRightsError{}
)
