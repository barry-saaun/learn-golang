package dictionary

type Dictionary map[string]string

var (
	ErrNotFound               = DictionaryError("could not found the word you were looking for")
	ErrWordExists             = DictionaryError("cannot add word because it already exists")
	ErrWordDoesNotExistUpdate = DictionaryError("word does not exist to be updated")
	ErrWordDoesNotExistDelete = DictionaryError("word does not exist to be deleted")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExistUpdate
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExistDelete
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
