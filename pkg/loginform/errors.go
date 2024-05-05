package loginform

import "fmt"

type IncorrectLoginOrPasswordError struct {
	login string
	pass  string
}

func newIncorrectLoginOrPasswordError(login, pass string) IncorrectLoginOrPasswordError {
	return IncorrectLoginOrPasswordError{
		login: login,
		pass:  pass,
	}
}

func (err IncorrectLoginOrPasswordError) Error() string {
	return fmt.Sprintf("incorrect login '%s' or password '%s'", err.login, err.pass)
}

type ExistLoginFormError struct {
	login string
}

func newExistLoginFormError(login string) ExistLoginFormError {
	return ExistLoginFormError{
		login: login,
	}
}

func (err ExistLoginFormError) Error() string {
	return fmt.Sprintf("user with login '%s' already exists", err.login)
}

var (
	IncorrectLoginOrPasswordErr = IncorrectLoginOrPasswordError{}
	ExistLoginFormErr           = ExistLoginFormError{}
)
