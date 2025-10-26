package blogposts_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
	"reflect"

	blogposts "github.com/RVSNS/blogposts"
)

type StubFailingFS struct{

}

func (fs StubFailingFS) Open(name string)(fs.File, error){
	return nil, errors.New("Always Fail!")
}

func assertEqual[T comparable](t *testing.T, got, want T){
	t.Helper()
	if got!=want{
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func assertPost(t *testing.T, got, want blogposts.Post){
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}
func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()

	if got.Title != want.Title {
		t.Fatalf("Title: got %q, want %q", got.Title, want.Title)
	}
	if got.Description != want.Description {
		t.Fatalf("Description: got %q, want %q", got.Description, want.Description)
	}

	if len(got.Tags) != len(want.Tags) {
		t.Fatalf("Tags: got %v, want %v", got.Tags, want.Tags)
	}
	for i := range got.Tags {
		if got.Tags[i] != want.Tags[i] {
			t.Fatalf("Tags[%d]: got %q, want %q", i, got.Tags[i], want.Tags[i])
		}
	}

	if got.Body != want.Body {
		t.Fatalf("Body: got %q, want %q", got.Body, want.Body)
	}
}
func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts,_ := blogposts.NewPostsFromFS(fs)
		// rest of test code cut for brevity
	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}