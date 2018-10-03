package main

import "fmt"

func main() {
	//obj := Constructor(2)
	//obj.Put(2, 1)
	//obj.Put(3, 2)
	//fmt.Println(obj.Get(3))
	//fmt.Println(obj.Get(2))
	//obj.Put(4, 3)
	//fmt.Println(obj.Get(2))
	//fmt.Println(obj.Get(3))
	//fmt.Println(obj.Get(4))

	obj := Constructor(2)
	obj.Put(1,1) // ok
	obj.Put(2,2) // ok
	fmt.Println(obj.Get(1)) // return 1, push 1 to top
	obj.Put(3,3) // pop out 2, push 1 to tail
	fmt.Println(obj.Get(2)) // return -1
	obj.Put(4,4) // pop out 1, push 3 to tail
	fmt.Println(obj.Get(1)) // return -1
	fmt.Println(obj.Get(3)) // return 3
	fmt.Println(obj.Get(4)) // return 4
}

type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}

type LRUCache struct {
	capacity int
	data     map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity,
		make(map[int]*Node),
		nil,
		nil,
	}

	return lru
}

func (this *LRUCache) Get(key int) int {
	value, found := this.data[key]
	if !found {
		return -1
	}

	if value == this.head && value == this.tail {
		// Do nothing ?
	} else if value == this.head {
		// Do nothing, already in front
	} else if value == this.tail {
		// Bring to front, set new tail
		value.left.right = nil
		this.tail = value.left

		value.left = nil
		value.right = this.head
		value.right.left = value
		this.head = value
	} else {
		value.left.right, value.right.left = value.right, value.left

		value.left = nil
		value.right = this.head
		value.right.left = value
		this.head = value

	}

	return value.value
}

func (this *LRUCache) Put(key int, value int) {
	exists := this.Get(key)
	if exists != -1 {
		existing, found := this.data[key]
		if found {
			existing.value = value
			return
		}
	}

	node := &Node{
		key,
		value,
		nil,
		nil,
	}

	if this.head != nil {
		node.right = this.head
		this.head.left = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
	}

	this.data[key] = node

	if len(this.data) > this.capacity {
		oldest := this.tail
		if oldest != nil {
			this.tail = oldest.left
			oldest.left.right = nil
			delete(this.data, oldest.key)
		}
	}

}
