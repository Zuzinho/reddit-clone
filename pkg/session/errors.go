package session

// BadSignMethodError - ошибка неправильного ключа подписи токена
type BadSignMethodError struct {
}

// InvalidTokenPayloadError - ошибка неправильной полезной нагрузки токена
type InvalidTokenPayloadError struct {
}

// InvalidTokenError - ошибка неправильного токена
type InvalidTokenError struct {
}

// Error возвращает текстовое представление ошибки
func (BadSignMethodError) Error() string {
	return "bad sign method"
}

// Error возвращает текстовое представление ошибки
func (InvalidTokenPayloadError) Error() string {
	return "invalid token payload"
}

// Error возвращает текстовое представление ошибки
func (InvalidTokenError) Error() string {
	return "invalid token"
}

var (
	BadSignMethodErr       = BadSignMethodError{}
	InvalidTokenPayloadErr = InvalidTokenPayloadError{}
	InvalidTokenErr        = InvalidTokenError{}
)
