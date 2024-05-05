package vote

type Type int

const (
	DownVoteType Type = -1
	UnVoteType   Type = 0
	UpVoteType   Type = 1
)

type keyType struct {
	userID string
	postID string
}

func newKeyType(postID, userID string) keyType {
	return keyType{
		userID: userID,
		postID: postID,
	}
}
