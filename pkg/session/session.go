package session

import (
	"time"
)

type Session struct {
	UserID string
	Exp    time.Time
}

func NewSession(userID string) *Session {
	return &Session{
		UserID: userID,
		Exp:    time.Now().Add(72 * time.Hour),
	}
}

func NewSessionWithExp(userID string, exp time.Time) *Session {
	return &Session{
		UserID: userID,
		Exp:    exp,
	}
}
