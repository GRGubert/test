package main

import (
	"28/pkg/post"
	"28/pkg/repo/post_repo"
	"28/pkg/repo/user_repo"
	"28/pkg/user"
	"fmt"
)

func main() {
	us := user_repo.NewUserStorage()
	ps := post_repo.NewPostStorage()

	tony := user.NewUser("Anthony")
	jhonny := user.NewUser("Jhonny Sack")
	tony.MakeFriend(jhonny)
	p := post.NewPost(tony.Name, "some post content")
	ps.Put(p)
	us.Put(tony)
	us.Put(jhonny)
	for _, users := range us {
		fmt.Println(users)
	}

	for _, post := range ps {
		fmt.Println(post)
	}
}
