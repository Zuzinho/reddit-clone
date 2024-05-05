package handlers

type NoAuthTokenError struct {
}

func (err NoAuthTokenError) Error() string {
	return "no access token"
}

var (
	NoAuthTokenErr = NoAuthTokenError{}
)
