package main

import (
	lcache "assignment-2/cache/lru"
	"fmt"
)

func main() {
	obj := lcache.CreateNewLRUCache(3)
	obj.Insert(1, 10)
	obj.Insert(3, 15)
	obj.Insert(2, 12)
	fmt.Println(obj.Get(3))
	obj.Insert(4, 20)
	fmt.Println(obj.Get(2))
	obj.Insert(4, 25)
	fmt.Println(obj.Get(6))
	// obj1 := fcache.CreateNewFifoCache(2)
	// obj1.Insert(1, 10)
	// obj1.Insert(3, 15)
	// obj1.Insert(2, 12)
	// fmt.Println(obj1.Get(3))
	// obj1.Insert(4, 20)
	// fmt.Println(obj1.Get(2))
	// obj1.Insert(4, 25)
	// fmt.Println(obj1.Get(6))
}
