package cache

type ICache interface {
	Set(key string, value int)
	Get(key string) int
	Delete(key string)
}

type d struct {
	data map[string]int
}

func New() *d {
	return &d{
		data: make(map[string]int),
	}
}

func (t *d) Set(key string, value int) {
	t.data[key] = value
}

func (t *d) Get(key string) int {
	return t.data[key]
}

func (t *d) Delete(key string) {
	delete(t.data, key)
}
