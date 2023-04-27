package main

import "fmt"

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")
	fmt.Println(cache.storage)

	cache.add("c", "3")
	fmt.Println(cache.storage)
	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")
	fmt.Println(cache.storage)

}
