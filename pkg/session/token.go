package session

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecretKey = []byte("Zuzinho secret key")

// PackToken запаковывает Session в токен
func PackToken(s *Session) (string, error) {
	payload := jwt.MapClaims{
		"sub": s.UserID,
		"iat": s.Exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// UnpackToken развертывает токен в Session
func UnpackToken(tokenString string) (*Session, error) {
	hashSecretGetter := func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method.Alg() != "HS256" {
			return nil, BadSignMethodErr
		}
		return jwtSecretKey, nil
	}

	token, err := jwt.Parse(tokenString, hashSecretGetter)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, InvalidTokenErr
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, InvalidTokenPayloadErr
	}

	userID := payload["sub"].(string)
	exp, err := time.Parse(time.RFC3339Nano, payload["iat"].(string))

	if err != nil {
		return nil, err
	}

	return NewSessionWithExp(userID, exp), nil
}
