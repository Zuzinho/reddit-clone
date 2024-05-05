package session

type BadSignMethodError struct {
}

type InvalidTokenPayloadError struct {
}

type InvalidTokenError struct {
}

func (BadSignMethodError) Error() string {
	return "bad sign method"
}

func (InvalidTokenPayloadError) Error() string {
	return "invalid token payload"
}

func (InvalidTokenError) Error() string {
	return "invalid token"
}

var (
	BadSignMethodErr       = BadSignMethodError{}
	InvalidTokenPayloadErr = InvalidTokenPayloadError{}
	InvalidTokenErr        = InvalidTokenError{}
)
