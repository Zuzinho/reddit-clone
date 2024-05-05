package handlers

import (
	"encoding/json"
	"io"
	"main/pkg/loginform"
	"main/pkg/session"
	"main/pkg/user"
	"net/http"
)

// AuthHandler - обработчик для регистрации
type AuthHandler struct {
	UsersRepo      user.UsersRepo
	LoginFormsRepo loginform.LoginFormsRepo
}

// NewAuthHandler создает экземпляр AuthHandler
func NewAuthHandler(usersRepo user.UsersRepo, loginFormsRepo loginform.LoginFormsRepo) *AuthHandler {
	return &AuthHandler{
		UsersRepo:      usersRepo,
		LoginFormsRepo: loginFormsRepo,
	}
}

// Register обрабатывает запрос на регистрации
func (handler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	f, err := handler.getBodyType(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.LoginFormsRepo.SignUp(f.Username, f.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := handler.UsersRepo.Create(f.Username)

	s := session.NewSession(u.ID)

	token, err := session.PackToken(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := authResponse{
		Token: token,
	}

	buf, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Login обрабатывает запрос на авторизацию
func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	f, err := handler.getBodyType(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.LoginFormsRepo.SignIn(f.Username, f.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := handler.UsersRepo.GetByUserName(f.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := session.NewSession(u.ID)

	token, err := session.PackToken(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := authResponse{
		Token: token,
	}

	buf, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *AuthHandler) getBodyType(r *http.Request) (*authForm, error) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return nil, err
	}

	var f authForm
	err = json.Unmarshal(body, &f)
	if err != nil {
		return nil, err
	}

	return &f, nil
}
