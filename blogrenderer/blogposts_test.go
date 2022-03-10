package blogrenderer_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	blogposts "github.com/umlx5h/learn-go-with-tests/blogrenderer"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World
`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
K
M
R`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	t.Run("post counts", func(t *testing.T) {
		posts, err := blogposts.NewPostsFromFS(fs)
		require.NoError(t, err)
		assert.Equal(t, len(fs), len(posts), "read from folder and count post num")
	})

	t.Run("error post created", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})
		assert.Error(t, err)
	})

	t.Run("post created", func(t *testing.T) {
		posts, err := blogposts.NewPostsFromFS(fs)
		require.NoError(t, err)

		got := posts[0]
		want := blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body:        "Hello\nWorld"}

		assert.Equal(t, want, got)
	})

}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
