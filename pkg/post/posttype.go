package post

// Type - тип поста
type Type string

const (
	TextType Type = "text"
	LinkType Type = "link"
)

// Valid проверяет валидность Type
func (tp Type) Valid() error {
	switch tp {
	case TextType, LinkType:
		return nil
	default:
		return newUnknownTypeError(tp)
	}
}
