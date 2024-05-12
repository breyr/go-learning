package dictionary

import "errors"

var ErrNotFound = errors.New("not found")

// a map value is a pointer to a runtime.hmap structure
// passing a map into a function/method you are copying the pointer, not the underlying data structure
type Dict map[string]string

func (d Dict) Search(key string) (string, error) {
	def, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dict) Add(word, def string) {
	d[word] = def
}
