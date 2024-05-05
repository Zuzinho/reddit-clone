package vote

// Vote - тип голоса
type Vote struct {
	UserID   string `json:"user"`
	postID   string
	VoteType Type `json:"vote"`
}

// NewVote возвращает экземпляр Vote
func NewVote(postID, userID string, voteType Type) *Vote {
	return &Vote{
		UserID:   userID,
		postID:   postID,
		VoteType: voteType,
	}
}

// Votes - массив ссылок на Vote
type Votes []*Vote

// Append добавляет Vote в Votes
func (votes *Votes) Append(vote *Vote) {
	*votes = append(*votes, vote)
}

// VotesRepo - интерфейс для хранения Vote
type VotesRepo interface {
	Create(postID, userID string, tp Type) *Vote
	Delete(postID, userID string)
	GetAllByPostID(postID string) *Votes
	DeleteByPostID(postID string)
}
