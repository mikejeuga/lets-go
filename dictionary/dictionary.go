package main

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFOund         = DictionaryErr("could not find a value for this key")
	ErrWordExists       = DictionaryErr("this word already exists in the dictionary")
	ErrWordDoesNotExist = DictionaryErr("This word is not in the dictionary")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFOund
	}
	return definition, nil
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFOund:
		d[key] = val
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, newDefintion string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFOund:
		return ErrWordDoesNotExist
	case nil:
		d[key] = newDefintion
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
