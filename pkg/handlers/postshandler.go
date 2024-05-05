package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"main/pkg/comment"
	"main/pkg/post"
	"main/pkg/user"
	"main/pkg/vote"
	"net/http"
	"strconv"
)

type PostsHandler struct {
	PostsRepo    post.PostsRepo
	UsersRepo    user.UsersRepo
	CommentsRepo comment.CommentsRepo
	VotesRepo    vote.VotesRepo
}

func NewPostsHandler(postsRepo post.PostsRepo, usersRepo user.UsersRepo, commentsRepo comment.CommentsRepo, votesRepo vote.VotesRepo) *PostsHandler {
	return &PostsHandler{
		PostsRepo:    postsRepo,
		UsersRepo:    usersRepo,
		CommentsRepo: commentsRepo,
		VotesRepo:    votesRepo,
	}
}

func (handler *PostsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts := handler.PostsRepo.GetAll()

	jsonPosts := handler.convertJSONPosts(posts)

	buf, err := json.Marshal(jsonPosts)
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

func (handler *PostsHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	var f createPostForm

	err := handler.getFormBody(r, &f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authorID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, NoAuthTokenErr.Error(), http.StatusBadRequest)
		return
	}

	u, err := handler.UsersRepo.GetByID(authorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var content string
	switch f.Type {
	case post.LinkType:
		content = f.URL
	case post.TextType:
		content = f.Text
	}

	p, err := handler.PostsRepo.Create(u, f.Title, content, f.Category, f.Type)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buf, err := json.Marshal(*handler.newJSONPostType(p))
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

	w.WriteHeader(http.StatusCreated)
}

func (handler *PostsHandler) GetAllByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category_name"]

	posts, err := handler.PostsRepo.GetAllByCategory(post.Category(category))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonPosts := handler.convertJSONPosts(posts)

	buf, err := json.Marshal(jsonPosts)
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

func (handler *PostsHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postID := vars["post_id"]

	p, err := handler.PostsRepo.Get(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.Views++

	jsonPost := handler.newJSONPostType(p)

	userID := r.Context().Value("user_id")
	userIDStr, ok := userID.(string)
	if ok {
		v := visitor{
			Admin: false,
			ID:    userIDStr,
		}

		jsonPost.User = &v
	}

	buf, err := json.Marshal(*jsonPost)
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

func (handler *PostsHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postID := vars["post_id"]

	var f createCommentForm
	err := handler.getFormBody(r, &f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authorID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, NoAuthTokenErr.Error(), http.StatusBadRequest)
		return
	}

	u, err := handler.UsersRepo.GetByID(authorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := handler.CommentsRepo.Create(postID, f.Comment, u)

	buf, err := json.Marshal(*c)
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

	w.WriteHeader(http.StatusCreated)
}

func (handler *PostsHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postID := vars["post_id"]

	commentID := vars["comment_id"]

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, NoAuthTokenErr.Error(), http.StatusBadRequest)
		return
	}

	err := handler.CommentsRepo.Delete(postID, commentID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *PostsHandler) Vote(voteType vote.Type) http.Handler {
	var path string

	switch voteType {
	case vote.UpVoteType:
		path = "upvote"
	case vote.UnVoteType:
		path = "unvote"
	case vote.DownVoteType:
		path = "downvote"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		postID := vars["post_id"]

		userID, ok := r.Context().Value("user_id").(string)
		if !ok {
			http.Error(w, NoAuthTokenErr.Error(), http.StatusBadRequest)
			return
		}

		switch path {
		case "upvote":
			handler.VotesRepo.Create(postID, userID, vote.UpVoteType)
		case "downvote":
			handler.VotesRepo.Create(postID, userID, vote.DownVoteType)
		case "unvote":
			handler.VotesRepo.Delete(postID, userID)
		}

		p, err := handler.PostsRepo.Get(postID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonPost := handler.newJSONPostType(p)

		buf, err := json.Marshal(jsonPost)
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
	})
}

func (handler *PostsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postID := vars["post_id"]

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, NoAuthTokenErr.Error(), http.StatusBadRequest)
		return
	}

	err := handler.PostsRepo.Delete(postID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	handler.CommentsRepo.DeleteByPostID(postID)
	handler.VotesRepo.DeleteByPostID(postID)

	w.WriteHeader(http.StatusOK)
}

func (handler *PostsHandler) GetAllByUserName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["user_login"]

	posts := handler.PostsRepo.GetAllByUserName(userName)

	jsonPosts := handler.convertJSONPosts(posts)

	buf, err := json.Marshal(jsonPosts)
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

func (handler *PostsHandler) uint32FromMap(vars map[string]string, key string) (uint32, error) {
	value := vars[key]

	d, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return uint32(d), nil
}
