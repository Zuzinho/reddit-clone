package handlers

// NoAuthTokenError - ошибка отсутствия токена авторизации
type NoAuthTokenError struct {
}

// Error возвращает текстовое описание ошибки
func (err NoAuthTokenError) Error() string {
	return "no access token"
}

var (
	NoAuthTokenErr = NoAuthTokenError{}
)
