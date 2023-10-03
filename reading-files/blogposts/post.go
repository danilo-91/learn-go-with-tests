package blogposts

import (
	"io"
	"io/fs"
)

func getPost(fsys fs.FS, name string) (Post, error) {
	file, err := fsys.Open(name)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()
	return readPost(file)
}

func readPost(f fs.File) (Post, error) {
	data, err := io.ReadAll(f)
	if err != nil {
		return Post{}, err
	}
	post := Post{Title: string(data)[7:]}
	return post, nil
}
