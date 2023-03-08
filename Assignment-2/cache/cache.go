package cache

//Cache interface --lower label of abstraction
type Cache interface {
	Get(key int) interface{}
	Insert(key int, value interface{})
}
