package cache

type Cache interface {
	Get(key int) interface{}
	Insert(key int, value interface{})
}
