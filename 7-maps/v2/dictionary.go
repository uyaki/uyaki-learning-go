package v1

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (dict Dictionary) Search(word string) (string, error) {
	// 第二个值是一个布尔值，表示是否成功找到 key。
	definition, ok := dict[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
