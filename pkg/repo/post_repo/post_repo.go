package post_repo

import (
	"28/pkg/post"
	"fmt"
)

type PostStorage map[string]*post.Post

func NewPostStorage() PostStorage {
	return make(map[string]*post.Post)
}

func (ps PostStorage) Put(p *post.Post) {
	ps[p.Author] = p
}

func (ps PostStorage) Get(authorName string) (*post.Post, error) {
	u, ok := ps[authorName]
	if !ok {
		return nil, fmt.Errorf("No such user")
	}
	return u, nil
}
