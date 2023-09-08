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
	_, err := d.Search(word)
	switch err {
	case nil:
		return ErrWordExists
	case ErrWordNotFound:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExists
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

const (
	ErrWordNotFound      = ErrDictionary("Error: word not found!")
	ErrWordExists        = ErrDictionary("Error: word already exists!")
	ErrWordDoesNotExists = ErrDictionary("Error: can't update/delete because word does not exists!")
)
