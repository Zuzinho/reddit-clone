package user

import (
	"sync"
)

// UsersMemoryRepository реализует интерфейс UsersRepo
type UsersMemoryRepository struct {
	mu   *sync.RWMutex
	data map[string]*User
}

// NewUsersMemoryRepository возвращает экземпляр UsersMemoryRepository
func NewUsersMemoryRepository() *UsersMemoryRepository {
	return &UsersMemoryRepository{
		mu:   &sync.RWMutex{},
		data: make(map[string]*User),
	}
}

// Create создает и добавляет User
func (repo *UsersMemoryRepository) Create(userName string) *User {
	repo.mu.Lock()

	u := NewUser(userName)

	repo.data[u.ID] = u

	repo.mu.Unlock()

	return u
}

// GetByID возвращает User c ID
func (repo *UsersMemoryRepository) GetByID(userID string) (*User, error) {
	repo.mu.RLock()
	u, ok := repo.data[userID]
	repo.mu.RUnlock()

	if !ok {
		return nil, newNoUserError(userID)
	}

	return u, nil
}

// GetByUserName возвращает User с UserName
func (repo *UsersMemoryRepository) GetByUserName(userName string) (*User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	for _, v := range repo.data {
		if v.UserName == userName {
			return v, nil
		}
	}

	return nil, newNoUserByUserNameError(userName)
}
