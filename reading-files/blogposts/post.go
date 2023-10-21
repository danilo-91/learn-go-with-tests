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

	return Post{
        Title: readLine(tPrefix, scanner),
        Description: readLine(dPrefix, scanner),
        Tags: readTags(tagsPrefix, scanner),
        Body: readBody(scanner),
    }, nil
}

// Use &scanner to extract text from file, trim prefix an return string
func readLine(prefix string, sc *bufio.Scanner) string {
    sc.Scan()
    return strings.TrimPrefix(sc.Text(), prefix)
}

// Use &scanner to extract text, trim prefix and split to []string
func readTags(prefix string, sc *bufio.Scanner) []string {
    sc.Scan()
    return strings.Split(strings.TrimPrefix(sc.Text(), prefix), ", ")
}

// Use &scanner to skip "---" and extract entire body
func readBody(sc *bufio.Scanner) string {
    sc.Scan() // skip "---"

    b := bytes.Buffer{}
    for sc.Scan() {
        fmt.Fprintln(&b, sc.Text())
    }
    return strings.TrimSuffix(b.String(), "\n")
}
