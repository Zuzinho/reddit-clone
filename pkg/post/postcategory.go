package post

type Category string

const (
	MusicCategory       Category = "music"
	FunnyCategory       Category = "funny"
	VideosCategory      Category = "videos"
	ProgrammingCategory Category = "programming"
	NewsCategory        Category = "news"
	FashionCategory     Category = "fashion"
)

func (category Category) Valid() error {
	switch category {
	case MusicCategory, FunnyCategory, VideosCategory, ProgrammingCategory, NewsCategory, FashionCategory:
		return nil
	default:
		return newUnknownCategoryError(category)
	}
}
