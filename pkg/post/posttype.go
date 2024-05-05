package post

type Type string

const (
	TextType Type = "text"
	LinkType Type = "link"
)

func (tp Type) Valid() error {
	switch tp {
	case TextType, LinkType:
		return nil
	default:
		return newUnknownTypeError(tp)
	}
}
