package v1

type Dictionary map[string]string

func (dict Dictionary) Search(key string) string {
	return dict[key]
}
