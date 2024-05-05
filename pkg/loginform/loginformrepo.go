package loginform

import (
	"sync"
)

type LoginFormsMemoryRepository struct {
	mu   *sync.RWMutex
	data map[string]*LoginForm
}

func NewLoginFormsMemoryRepository() *LoginFormsMemoryRepository {
	return &LoginFormsMemoryRepository{
		mu:   &sync.RWMutex{},
		data: make(map[string]*LoginForm),
	}
}

func (repo *LoginFormsMemoryRepository) SignIn(login, pass string) error {
	repo.mu.RLock()
	form, ok := repo.data[login]
	repo.mu.RUnlock()

	if !ok {
		return newIncorrectLoginOrPasswordError(login, pass)
	}

	if !form.CheckPassword(pass) {
		return newIncorrectLoginOrPasswordError(login, pass)
	}

	return nil
}

func (repo *LoginFormsMemoryRepository) SignUp(login, pass string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, ok := repo.data[login]; ok {
		return newExistLoginFormError(login)
	}

	form := NewLoginForm(login, pass)
	repo.data[login] = form

	return nil
}
