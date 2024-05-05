package vote

import "sync"

type VotesMemoryRepository struct {
	mu    *sync.RWMutex
	votes map[keyType]*Vote
}

func NewVotesMemoryRepository() *VotesMemoryRepository {
	return &VotesMemoryRepository{
		mu:    &sync.RWMutex{},
		votes: make(map[keyType]*Vote),
	}
}

func (repo *VotesMemoryRepository) Create(postID, userID string, tp Type) *Vote {
	k := newKeyType(postID, userID)
	v := NewVote(postID, userID, tp)

	repo.mu.Lock()
	repo.votes[k] = v
	repo.mu.Unlock()

	return v
}

func (repo *VotesMemoryRepository) Delete(postID, userID string) {
	k := newKeyType(postID, userID)

	repo.mu.Lock()
	delete(repo.votes, k)
	repo.mu.Unlock()
}

func (repo *VotesMemoryRepository) GetAllByPostID(postID string) *Votes {
	votes := make(Votes, 0)

	repo.mu.RLock()
	for k, v := range repo.votes {
		if k.postID == postID {
			votes.Append(v)
		}
	}
	repo.mu.RUnlock()

	return &votes
}

func (repo *VotesMemoryRepository) DeleteByPostID(postID string) {
	repo.mu.Lock()
	for k := range repo.votes {
		if k.postID == postID {
			delete(repo.votes, k)
		}
	}
	repo.mu.RLock()
}
