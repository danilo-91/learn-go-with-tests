package blogposts_test

import (
	"errors"
	"github.com/isedaniel/blogposts"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("2 files", func(t *testing.T) {
		const (
			post1Body = `Title: Post 1
Description: Description 1`
			post2Body = `Title: Post 2
Description: Description 2`
		)

		fsys := fstest.MapFS{
			"hello world.md":  {Data: []byte(post1Body)},
			"hello-world2.md": {Data: []byte(post2Body)},
		}

		posts, err := blogposts.NewPostsFromFS(fsys)
		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		want := blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
		}
		assertPosts(t, got, want)
	})

	t.Run("error handling", func(t *testing.T) {
		fs := failingFS{}

		_, err := blogposts.NewPostsFromFS(fs)

		if err == nil {
			t.Errorf("Expected error, got %v", err)
		}
	})
}

type failingFS struct{}

func (fs failingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Sry, I'll allways fail.")
}

func assertPosts(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, but wanted %+v", got, want)
	}
}
