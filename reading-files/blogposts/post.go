package blogposts

import (
    "bufio"
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

const (
    titleSeparator = "Title: "
    descSeparator = "Description: "
)

func readPost(f fs.File) (Post, error) {
    scanner := bufio.NewScanner(f)

    readLine := func() string {
        scanner.Scan()
        return scanner.Text()
    }

    title := readLine()
    description := readLine()

	post := Post{
        Title: title[len(titleSeparator):],
        Description: description[len(descSeparator):],
    }
	return post, nil
}
