package post

import (
	"main/pkg/user"
	"sync"
)

type PostsMemoryRepository struct {
	mu    *sync.RWMutex
	posts map[string]*Post
}

func NewPostsMemoryRepository() *PostsMemoryRepository {
	return &PostsMemoryRepository{
		mu:    &sync.RWMutex{},
		posts: make(map[string]*Post),
	}
}

func (repo *PostsMemoryRepository) Create(author *user.User, title, text string, category Category, tp Type) (*Post, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	p, err := NewPost(author, title, text, category, tp)
	if err != nil {
		return nil, err
	}

	repo.posts[p.ID] = p

	return p, nil
}

func (repo *PostsMemoryRepository) Get(postID string) (*Post, error) {
	repo.mu.RLock()
	p, ok := repo.posts[postID]
	repo.mu.RUnlock()

	if !ok {
		return nil, newNoPostError(postID)
	}

	return p, nil
}

func (repo *PostsMemoryRepository) filter(condition func(post *Post) bool) *Posts {
	posts := make(Posts, 0)

	repo.mu.RLock()
	for _, v := range repo.posts {
		if condition(v) {
			posts = append(posts, v)
		}
	}
	repo.mu.RUnlock()

	return &posts
}

func (repo *PostsMemoryRepository) GetAll() *Posts {
	return repo.filter(func(*Post) bool {
		return true
	})
}

func (repo *PostsMemoryRepository) GetAllByUserName(userName string) *Posts {
	return repo.filter(func(post *Post) bool {
		return post.Author.UserName == userName
	})
}

func (repo *PostsMemoryRepository) GetAllByCategory(category Category) (*Posts, error) {
	err := category.Valid()
	if err != nil {
		return nil, err
	}

	return repo.filter(func(post *Post) bool {
		return post.Category == category
	}), nil
}

func (repo *PostsMemoryRepository) Delete(postID, userID string) error {
	repo.mu.RLock()
	p, ok := repo.posts[postID]
	repo.mu.RUnlock()

	if !ok {
		return newNoPostError(postID)
	}

	if p.Author.ID != userID {
		return newNoRightsError(postID)
	}

	repo.mu.Lock()
	delete(repo.posts, postID)
	repo.mu.RUnlock()

	return nil
}
