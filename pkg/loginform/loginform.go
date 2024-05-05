package loginform

// LoginForm - форма авторизации
type LoginForm struct {
	UserName string
	password string
}

// NewLoginForm возвращает экземпляр LoginForm
func NewLoginForm(userName, password string) *LoginForm {
	return &LoginForm{
		UserName: userName,
		password: password,
	}
}

// CheckPassword проверяет пароль для LoginForm
func (form *LoginForm) CheckPassword(password string) bool {
	return form.password == password
}

// LoginFormsRepo - интерфейс для хранения LoginForm
type LoginFormsRepo interface {
	SignIn(login, pass string) error
	SignUp(login, pass string) error
}
