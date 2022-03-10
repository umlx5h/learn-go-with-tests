package blogrenderer_test

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/stretchr/testify/require"
	"github.com/umlx5h/learn-go-with-tests/blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)
		require.NoError(t, err)

		approvals.VerifyString(t, buf.String())

		// // got := buf.String()
		// got := strings.ReplaceAll(buf.String(), "\n", "")
		// want := `<h1>hello world</h1><p>This is a description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`
		// assert.Equal(t, want, got)
	})
}
