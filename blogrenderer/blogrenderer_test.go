package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/stretchr/testify/require"
	"github.com/umlx5h/learn-go-with-tests/blogrenderer"
)

var (
	aPost = blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
)

func TestRender(t *testing.T) {
	postRenderer, err := blogrenderer.NewPostRenderer()
	require.NoError(t, err)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err = postRenderer.Render(&buf, aPost)
		require.NoError(t, err)

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{
			{Title: "Hello World"},
			{Title: "Hello World 2"},
		}

		err := postRenderer.RenderIndex(&buf, posts)
		require.NoError(t, err)

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	postRenderer, err := blogrenderer.NewPostRenderer()
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
