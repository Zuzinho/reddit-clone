package post

import "fmt"

// UnknownCategoryError - ошибка неизвестной категории Post
type UnknownCategoryError struct {
	category Category
}

func newUnknownCategoryError(category Category) UnknownCategoryError {
	return UnknownCategoryError{
		category: category,
	}
}

// Error возвращает текстовое представление ошибки
func (err UnknownCategoryError) Error() string {
	return fmt.Sprintf("unknown category '%s'", err.category)
}

// UnknownPostTypeError - ошибка неизвестного типа Post
type UnknownPostTypeError struct {
	postType Type
}

func newUnknownTypeError(postType Type) UnknownPostTypeError {
	return UnknownPostTypeError{
		postType: postType,
	}
}

// Error возвращает текстовое представление ошибки
func (err UnknownPostTypeError) Error() string {
	return fmt.Sprintf("unknown post type '%s'", err.postType)
}

// NoPostError - ошибка отсутствия Post
type NoPostError struct {
	postID string
}

func newNoPostError(postID string) NoPostError {
	return NoPostError{
		postID: postID,
	}
}

// Error возвращает текстовое представление ошибки
func (err NoPostError) Error() string {
	return fmt.Sprintf("no post by id '%s'", err.postID)
}

// NoRightsError - ошибка отсутствия прав
type NoRightsError struct {
	postID string
}

func newNoRightsError(postID string) NoRightsError {
	return NoRightsError{
		postID: postID,
	}
}

// Error возвращает текстовое представление ошибки
func (err NoRightsError) Error() string {
	return fmt.Sprintf("have no rights for post '%s'", err.postID)
}

var (
	UnknownCategoryErr = UnknownCategoryError{}
	UnknownPostTypeErr = UnknownPostTypeError{}
	NoPostErr          = NoPostError{}
	NoRightsErr        = NoRightsError{}
)
