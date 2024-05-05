package handlers

import (
	"encoding/json"
	"io"
	"main/pkg/comment"
	"main/pkg/post"
	"main/pkg/vote"
	"net/http"
	"sync"
)

type visitor struct {
	ID    string `json:"id"`
	Admin bool   `json:"admin"`
}

type jsonPostType struct {
	post.Post
	Score            int               `json:"score"`
	UpvotePercentage float32           `json:"upvotePercentage"`
	Comments         *comment.Comments `json:"comments"`
	Votes            *vote.Votes       `json:"votes"`
	User             *visitor          `json:"user"`
}

func (handler *PostsHandler) newJSONPostType(post *post.Post) *jsonPostType {
	wg := &sync.WaitGroup{}

	wg.Add(2)
	var comments *comment.Comments
	var votes *vote.Votes

	go func() {
		comments = handler.CommentsRepo.GetAllByPostID(post.ID)
		wg.Done()
	}()

	var score int

	go func() {
		votes = handler.VotesRepo.GetAllByPostID(post.ID)
		for _, v := range *votes {
			score += int(v.VoteType)
		}

		wg.Done()
	}()

	wg.Wait()

	var percent float32

	if post.Views > 0 {
		percent = float32(len(*votes)) / float32(post.Views)
	}

	return &jsonPostType{
		Post:             *post,
		Comments:         comments,
		Votes:            votes,
		UpvotePercentage: percent,
		Score:            score,
	}
}

func (handler *PostsHandler) convertJSONPosts(posts *post.Posts) []*jsonPostType {
	ln := len(*posts)

	jsonPosts := make([]*jsonPostType, ln)

	wg := &sync.WaitGroup{}
	wg.Add(ln)
	for i, p := range *posts {
		go func(i int, p *post.Post) {
			defer wg.Done()

			jsonPosts[i] = handler.newJSONPostType(p)
		}(i, p)
	}

	wg.Wait()

	return jsonPosts
}

type createPostForm struct {
	Type     post.Type     `json:"type"`
	Category post.Category `json:"category"`
	Title    string        `json:"title"`
	URL      string        `json:"url"`
	Text     string        `json:"text"`
}

type createCommentForm struct {
	Comment string
}

func (handler *PostsHandler) getFormBody(r *http.Request, f any) error {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &f)
	if err != nil {
		return err
	}

	return nil
}
