package dictionary

import "fmt"

type Dictionary map[string]string

func (d Dictionary) Search(s string) (string, error) {
	word, ok := d[s]

	if !ok {
		return "", ErrWordNotFound(s)
	}
	return word, nil
}

type ErrWordNotFound string

func (e ErrWordNotFound) Error() string {
	return fmt.Sprintf("Error: word %q not found!", string(e))
}
