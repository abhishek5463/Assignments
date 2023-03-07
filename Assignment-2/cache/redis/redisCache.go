package redis

import "assignment-2/cache"

type Redis struct{}

func NewRedis() cache.Cache {
	return &Redis{}
}

func (r *Redis) Get(key int) interface{} {
	return "haha"
}

func (r *Redis) Insert(key int, value interface{}) {
}
