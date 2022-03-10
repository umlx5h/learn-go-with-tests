package main

import (
	"log"
	"os"

	blogposts "github.com/umlx5h/learn-go-with-tests/blogrenderer"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", posts)
}
