package user

import "28/pkg/post"

type User struct {
	Name    string
	Post    []*post.Post
	Friends []*User
}

func NewUser(name string) *User {
	return &User{
		Name:    name,
		Post:    make([]*post.Post, 0),
		Friends: make([]*User, 0),
	}
}

func (u *User) MakeFriend(userFriend *User) {
	u.Friends = append(u.Friends, userFriend)
}

func (u *User) AddPost(p *post.Post) {
	u.Post = append(u.Post, p)
}
