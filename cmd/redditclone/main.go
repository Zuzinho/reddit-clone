package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"main/pkg/comment"
	"main/pkg/handlers"
	"main/pkg/loginform"
	"main/pkg/middleware"
	"main/pkg/post"
	"main/pkg/user"
	"main/pkg/vote"
	"net/http"
)

func main() {
	usersRepo := user.NewUsersMemoryRepository()
	loginFormsRepo := loginform.NewLoginFormsMemoryRepository()
	commentsRepo := comment.NewCommentsMemoryRepository()
	votesRepo := vote.NewVotesMemoryRepository()
	postsRepo := post.NewPostsMemoryRepository()
	appMiddleware := middleware.NewAppMiddleware()

	authHandler := handlers.NewAuthHandler(usersRepo, loginFormsRepo)
	postsHandler := handlers.NewPostsHandler(postsRepo, usersRepo, commentsRepo, votesRepo)

	commonRouter := mux.NewRouter()

	commonRouter.HandleFunc("/posts", postsHandler.AddPost).
		Methods("POST")
	commonRouter.HandleFunc("/post/{post_id:.+}", postsHandler.AddComment).
		Methods("POST")
	commonRouter.HandleFunc("/post/{post_id:.+}/{comment_id:.+}", postsHandler.DeleteComment).
		Methods("DELETE")
	commonRouter.Handle("/post/{post_id:.+}/upvote", postsHandler.Vote(vote.UpVoteType)).
		Methods("GET")
	commonRouter.Handle("/post/{post_id:.+}/downvote", postsHandler.Vote(vote.DownVoteType)).
		Methods("GET")
	commonRouter.Handle("/post/{post_id:.+}/unvote", postsHandler.Vote(vote.UnVoteType)).
		Methods("GET")
	commonRouter.HandleFunc("/post/{post_id:.+}", postsHandler.Delete).
		Methods("DELETE")

	commonRouter.HandleFunc("/register", authHandler.Register).
		Methods("POST")
	commonRouter.HandleFunc("/login", authHandler.Login).
		Methods("POST")

	commonRouter.HandleFunc("/posts/", postsHandler.GetAll).
		Methods("GET")

	commonRouter.HandleFunc("/posts/{category_name:[A-z]+}", postsHandler.GetAllByCategory).
		Methods("GET")
	commonRouter.HandleFunc("/post/{post_id:.+}", postsHandler.Get).
		Methods("GET")
	commonRouter.HandleFunc("/user/{user_login:[A-z0-9]+}", postsHandler.GetAllByUserName).
		Methods("GET")

	api := mux.NewRouter()

	api.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	api.PathPrefix("/api/").Handler(http.StripPrefix("/api", commonRouter))

	api.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./static/html/index.html"))
		err := tmpl.Execute(w, r)
		if err != nil {
			log.Println(err)
		}
	})

	log.Println("starting on port 8080")

	log.Println(http.ListenAndServe(":8080", appMiddleware.PackMiddleware(api)))
}
