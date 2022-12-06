package user_repo

import (
	"28/pkg/user"
	"fmt"
)

type UserStorage map[string]*user.User

func NewUserStorage() UserStorage {
	return make(map[string]*user.User)
}

func (us UserStorage) Put(u *user.User) {
	us[u.Name] = u
}

func (us UserStorage) Get(userName string) (*user.User, error) {
	u, ok := us[userName]
	if !ok {
		return nil, fmt.Errorf("No such user")
	}
	return u, nil
}
