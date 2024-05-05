package comment

type keyType struct {
	commentID string
	postID    string
}

func newKeyType(postID, commentID string) keyType {
	return keyType{
		postID:    postID,
		commentID: commentID,
	}
}
