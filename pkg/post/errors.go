package post

import "fmt"

type UnknownCategoryError struct {
	category Category
}

func newUnknownCategoryError(category Category) UnknownCategoryError {
	return UnknownCategoryError{
		category: category,
	}
}

func (err UnknownCategoryError) Error() string {
	return fmt.Sprintf("unknown category '%s'", err.category)
}

type UnknownPostTypeError struct {
	postType Type
}

func newUnknownTypeError(postType Type) UnknownPostTypeError {
	return UnknownPostTypeError{
		postType: postType,
	}
}

func (err UnknownPostTypeError) Error() string {
	return fmt.Sprintf("unknown post type '%s'", err.postType)
}

type NoPostError struct {
	postID string
}

func newNoPostError(postID string) NoPostError {
	return NoPostError{
		postID: postID,
	}
}

func (err NoPostError) Error() string {
	return fmt.Sprintf("no post by id '%s'", err.postID)
}

type NoRightsError struct {
	postID string
}

func newNoRightsError(postID string) NoRightsError {
	return NoRightsError{
		postID: postID,
	}
}

func (err NoRightsError) Error() string {
	return fmt.Sprintf("have no rights for post '%s'", err.postID)
}

var (
	UnknownCategoryErr = UnknownCategoryError{}
	UnknownPostTypeErr = UnknownPostTypeError{}
	NoPostErr          = NoPostError{}
	NoRightsErr        = NoRightsError{}
)
