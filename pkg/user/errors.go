package user

import (
	"fmt"
)

type NoUserError struct {
	id string
}

func newNoUserError(id string) NoUserError {
	return NoUserError{
		id: id,
	}
}

func (err NoUserError) Error() string {
	return fmt.Sprintf("no user by id '%s'", err.id)
}

type NoUserByUserNameError struct {
	userName string
}

func newNoUserByUserNameError(userName string) NoUserByUserNameError {
	return NoUserByUserNameError{
		userName: userName,
	}
}

func (err NoUserByUserNameError) Error() string {
	return fmt.Sprintf("user with username '%s' already exist", err.userName)
}

var (
	NoUserErr           = NoUserError{}
	NoUserByUserNameErr = NoUserByUserNameError{}
)
