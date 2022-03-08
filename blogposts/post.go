package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	getTag := func(tagName string) []string {
		tagText := readMetaLine(tagSeparator)
		tags := strings.Split(tagText, ",")

		for i := range tags {
			tags[i] = strings.TrimSpace(tags[i])
		}

		return tags
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tag := getTag(tagSeparator)

	return Post{
		Title:       title,
		Description: description,
		Tags:        tag,
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore ---

	var buf bytes.Buffer
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	body := strings.TrimSuffix(buf.String(), "\n")

	return body
}
