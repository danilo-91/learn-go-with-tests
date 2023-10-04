package blogposts

import (
	"bufio"
	"io/fs"
	"strings"
    "fmt"
    "bytes"
)

// Open post file using name from fileSystem fsys
func getPost(fsys fs.FS, name string) (Post, error) {
	file, err := fsys.Open(name)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()
	return readPost(file)
}

const (
	tPrefix = "Title: "
	dPrefix  = "Description: "
    tagsPrefix = "Tags: "
)

// Scan file data to fetch post
func readPost(f fs.File) (Post, error) {
	scanner := bufio.NewScanner(f)

	readLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

    title := readLine(tPrefix)
    desc := readLine(dPrefix)
    tags := strings.Split(readLine(tagsPrefix), ", ")

    readLine("") // skip line of "---"

    buf := bytes.Buffer{}
    for scanner.Scan() {
        fmt.Fprintln(&buf, scanner.Text())
    }
    body := strings.TrimSuffix(buf.String(), "\n")

	return Post{
        Title: title,
        Description: desc,
        Tags: tags,
        Body: body,
    }, nil
}
