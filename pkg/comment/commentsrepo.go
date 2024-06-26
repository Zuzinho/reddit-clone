package comment

import (
	"main/pkg/user"
	"sync"
)

// CommentsMemoryRepository реализует интерфейс CommentsRepo
type CommentsMemoryRepository struct {
	mu       *sync.RWMutex
	comments map[keyType]*Comment
}

// NewCommentsMemoryRepository создает CommentsMemoryRepository
func NewCommentsMemoryRepository() *CommentsMemoryRepository {
	return &CommentsMemoryRepository{
		mu:       &sync.RWMutex{},
		comments: make(map[keyType]*Comment),
	}
}

// Create создает и добавляет в базу Comment
func (repo *CommentsMemoryRepository) Create(postID string, body string, author *user.User) *Comment {
	repo.mu.Lock()

	c := NewComment(body, author)
	k := newKeyType(postID, c.ID)

	repo.comments[k] = c

	repo.mu.Unlock()

	return c
}

// Delete удаляет из базы Comment
func (repo *CommentsMemoryRepository) Delete(postID, commentID, userID string) error {
	k := newKeyType(postID, commentID)

	repo.mu.RLock()
	c, ok := repo.comments[k]
	repo.mu.RUnlock()

	if !ok {
		return newNoCommentError(commentID)
	}

	if c.Author.ID != userID {
		return newNoRightsError(commentID)
	}

	repo.mu.Lock()
	delete(repo.comments, k)
	repo.mu.Unlock()

	return nil
}

// GetAllByPostID возвращает Comments для поста
func (repo *CommentsMemoryRepository) GetAllByPostID(postID string) *Comments {
	comments := make(Comments, 0)

	repo.mu.RLock()
	for k, v := range repo.comments {
		if k.postID == postID {
			comments.Append(v)
		}
	}
	repo.mu.RUnlock()

	return &comments
}

// DeleteByPostID удаляет Comments для поста
func (repo *CommentsMemoryRepository) DeleteByPostID(postID string) {
	repo.mu.Lock()
	for k := range repo.comments {
		if k.postID == postID {
			delete(repo.comments, k)
		}
	}
	repo.mu.Unlock()
}
