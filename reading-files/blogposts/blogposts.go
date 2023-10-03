package blogposts

import (
	"io/fs"
)

type Post struct {
	Title, Description string
}

func NewPostsFromFS(fsys fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fsys, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, file := range dir {
		post, err := getPost(fsys, file.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
