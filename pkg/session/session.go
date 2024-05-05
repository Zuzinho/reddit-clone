package session

import (
	"time"
)

// Session - сессия пользователя
type Session struct {
	UserID string
	Exp    time.Time
}

// NewSession возвращает экземпляр Session
func NewSession(userID string) *Session {
	return &Session{
		UserID: userID,
		Exp:    time.Now().Add(72 * time.Hour),
	}
}

// NewSessionWithExp возвращает экземпляр со сроком годности
func NewSessionWithExp(userID string, exp time.Time) *Session {
	return &Session{
		UserID: userID,
		Exp:    exp,
	}
}
