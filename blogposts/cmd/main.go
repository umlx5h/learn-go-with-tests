package main

import (
	"log"
	"os"

	"github.com/umlx5h/learn-go-with-tests/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", posts)
}
