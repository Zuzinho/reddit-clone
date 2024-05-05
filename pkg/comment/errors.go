package comment

import "fmt"

type NoCommentError struct {
	commentID string
}

func newNoCommentError(commentID string) NoCommentError {
	return NoCommentError{
		commentID: commentID,
	}
}

func (err NoCommentError) Error() string {
	return fmt.Sprintf("no comment '%s'", err.commentID)
}

type NoRightsError struct {
	commentID string
}

func newNoRightsError(commentID string) NoRightsError {
	return NoRightsError{
		commentID: commentID,
	}
}

func (err NoRightsError) Error() string {
	return fmt.Sprintf("have no rights for comment '%s'", err.commentID)
}

var (
	NoCommentErr = NoCommentError{}
	NoRightsErr  = NoRightsError{}
)
