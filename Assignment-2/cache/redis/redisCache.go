package redis

import "assignment-2/cache"

type Redis struct{}

func NewRedis() cache.Cache {
	return &Redis{}
}

//another implementation of cache interface

func (r *Redis) Get(key int) interface{} {
	return -1
}

func (r *Redis) Insert(key int, value interface{}) {
}
