package dictionary

type Dictionary map[string]string

func (d Dictionary) Search(s string) (string, error) {
	word, ok := d[s]

	if !ok {
		return "", ErrWordNotFound
	}
	return word, nil
}

func (d Dictionary) Add(word, definition string) error {
    _, ok := d[word]
    if ok {
        return ErrWordExists
    }
	d[word] = definition
	return nil
}

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

const (
	ErrWordNotFound = ErrDictionary("Error: word not found!")
	ErrWordExists   = ErrDictionary("Error: word already exists!")
)
