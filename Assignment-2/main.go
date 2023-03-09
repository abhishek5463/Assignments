package main

import (
	fcache "assignment-2/cache/fifo"
	lcache "assignment-2/cache/lru"
	"fmt"
)

func main() {

	//LRU object creation
	fmt.Println("LRU")
	obj := lcache.CreateNewLRUCache(3)
	obj.Insert(1, 10)
	obj.Insert(3, 15)
	obj.Insert(2, 12)
	fmt.Println(obj.Get(3))
	obj.Insert(4, 20)
	fmt.Println(obj.Get(2))
	obj.Insert(4, 25)
	fmt.Println(obj.Get(6))

	fmt.Println("\nFIFO")
	//Fifo object creation

	obj1 := fcache.CreateNewFifoCache(2)
	obj1.Insert(10, 13)
	obj1.Insert(8, 5)
	obj1.Insert(9, 10)
	fmt.Println(obj1.Get(8))
	obj1.Insert(4, 20)
	fmt.Println(obj1.Get(10))
	obj1.Insert(4, 25)
	fmt.Println(obj1.Get(4))
}
