package loginform

type LoginForm struct {
	UserName string
	password string
}

func NewLoginForm(userName, password string) *LoginForm {
	return &LoginForm{
		UserName: userName,
		password: password,
	}
}

func (form *LoginForm) CheckPassword(password string) bool {
	return form.password == password
}

type LoginFormsRepo interface {
	SignIn(login, pass string) error
	SignUp(login, pass string) error
}
