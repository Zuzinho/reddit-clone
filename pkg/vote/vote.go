package vote

type Vote struct {
	UserID   string `json:"user"`
	postID   string
	VoteType Type `json:"vote"`
}

func NewVote(postID, userID string, voteType Type) *Vote {
	return &Vote{
		UserID:   userID,
		postID:   postID,
		VoteType: voteType,
	}
}

type Votes []*Vote

func (votes *Votes) Append(vote *Vote) {
	*votes = append(*votes, vote)
}

type VotesRepo interface {
	Create(postID, userID string, tp Type) *Vote
	Delete(postID, userID string)
	GetAllByPostID(postID string) *Votes
	DeleteByPostID(postID string)
}
