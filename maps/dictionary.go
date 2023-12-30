package maps

type Dictionary map[string]string

const (
	WordNotPresent     = DictionaryErr("word not present in the dictionary")
	wordAlreadyPresent = DictionaryErr("word already present in the dictionary")
	UpdateNewWord      = DictionaryErr("updating a new word")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return d.Error()
}

func (d Dictionary) Search(search string) (string, error) {
	definition, found := d[search]
	if !found {
		return "", WordNotPresent
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case WordNotPresent:
		d[word] = definition
	case nil:
		return wordAlreadyPresent

	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case WordNotPresent:
		return UpdateNewWord
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
