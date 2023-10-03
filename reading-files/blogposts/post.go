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

func readPost(f fs.File) (Post, error) {
    scanner := bufio.NewScanner(f)

    scanner.Scan()
    titleLine := scanner.Text()

    scanner.Scan()
    descriptionLine := scanner.Text()

	post := Post{
        Title: titleLine[7:],
        Description: descriptionLine[13:],
    }
	return post, nil
}
