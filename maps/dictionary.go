package main

type Dictionary map[string]string

// var (
// 	ErrNotFound   = errors.New("could not find the word you were looking for")
// 	ErrWordExists = errors.New("already exists word")
// )

type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("already exists word")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	if w, ok := d[word]; ok {
		return w, nil
	}

	return "", ErrNotFound
}

func (d Dictionary) Add(word, definition string) error {
	if _, err := d.Search(word); err == nil {
		return ErrWordExists
	}
	d[word] = definition

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
