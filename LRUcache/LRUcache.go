package main

import "fmt"

type Entry struct {
	key, value int // key, value >= 0
	prev, next *Entry
}
type LRUCache struct {
	dict       map[int]*Entry
	Capacity   int
	head, last *Entry // leftmost is LRU, rightmost is most recent

}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 {
		panic("LRUCache: capacity must be positive")
	}
	lru := LRUCache{
		Capacity: capacity,
		dict:     make(map[int]*Entry),
		head:     nil,
		last:     nil,
	}

	return lru
}

func (this *LRUCache) Get(key int) int {

	if entry, exists := this.dict[key]; exists {
		this.moveToFront(entry)
		return entry.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if entry, exists := this.dict[key]; exists {
		entry.value = value
		this.moveToFront(entry)
	} else {
		newNode := &Entry{
			key:   key,
			value: value,
		}
		this.dict[key] = newNode
		this.addToFront(newNode)
		if len(this.dict) > this.Capacity {
            this.evictLRU()
        }
	}

}

func (this *LRUCache) moveToFront(entry *Entry) {
	if entry == this.head {
		return
	}
	this.removeFromList(entry)
	this.addToFront(entry)

}

func (this *LRUCache) addToFront(entry *Entry) {
	entry.prev = nil
	entry.next = this.head

	if this.head != nil {
		entry.next.prev = entry
	}
	this.head = entry
	if this.last == nil {
		this.last = entry
	}
}

func (this *LRUCache) removeFromList(entry *Entry) {
	if entry.prev != nil {
		entry.prev.next = entry.next
	} else {
		this.head = entry.next
	}

	if entry.next != nil {
		entry.next.prev = entry.prev
	} else {
		this.last = entry.prev
	}
}

func (this *LRUCache) evictLRU() {
	if this.last != nil {
		delete(this.dict, this.last.key)
		this.removeFromList(this.last)
	}
}

/*


// ListNode represents a node in the doubly linked list
type ListNode1 struct {
    Key   int
    Value int
    Prev  *ListNode1
    Next  *ListNode1
}

// LRUCache represents the Least Recently Used (LRU) cache
type LRUCache struct {
    Capacity int
    Cache    map[int]*ListNode1
    Head     *ListNode1
    Tail     *ListNode1
}

// Constructor initializes the LRUCache with the given capacity
func Constructor(capacity int) LRUCache {
    return LRUCache{
        Capacity: capacity,
        Cache:    make(map[int]*ListNode1),
        Head:     nil,
        Tail:     nil,
    }
}

// Get returns the value of the key if it exists, otherwise returns -1
func (lru *LRUCache) Get(key int) int {
    if node, ok := lru.Cache[key]; ok {
        // Move the accessed node to the front of the list (most recently used)
        lru.moveToFront(node)
        return node.Value
    }
    return -1
}

// Put updates the value of the key if it exists. Otherwise, adds the key-value pair to the cache.
// If the number of keys exceeds the capacity, evicts the least recently used key.
func (lru *LRUCache) Put(key int, value int) {
    if node, ok := lru.Cache[key]; ok {
        // Update the value and move the node to the front (most recently used)
        node.Value = value
        lru.moveToFront(node)
    } else {
        // Add a new node to the cache
        newNode := &ListNode1{Key: key, Value: value}
        lru.Cache[key] = newNode

        // Add the new node to the front (most recently used)
        lru.addToFront(newNode)

        // Check if the cache exceeds its capacity and evict the least recently used node if needed
        if len(lru.Cache) > lru.Capacity {
            lru.evictLRU()
        }
    }
}

// moveToFront moves the specified node to the front of the linked list
func (lru *LRUCache) moveToFront(node *ListNode1) {
    // If the node is already at the front, do nothing
    if node == lru.Head {
        return
    }

    // Remove the node from its current position
    lru.removeFromList(node)

    // Add the node to the front
    lru.addToFront(node)
}

// addToFront adds the specified node to the front of the linked list
func (lru *LRUCache) addToFront(node *ListNode1) {
    node.Prev = nil
    node.Next = lru.Head

    if lru.Head != nil {
        lru.Head.Prev = node
    }

    lru.Head = node

    // If the list was empty, update the tail as well
    if lru.Tail == nil {
        lru.Tail = node
    }
}

// removeFromList removes the specified node from the linked list
func (lru *LRUCache) removeFromList(node *ListNode1) {
    if node.Prev != nil {
        node.Prev.Next = node.Next
    } else {
        lru.Head = node.Next
    }

    if node.Next != nil {
        node.Next.Prev = node.Prev
    } else {
        lru.Tail = node.Prev
    }
}

// evictLRU evicts the least recently used node from the cache
func (lru *LRUCache) evictLRU() {
    if lru.Tail != nil {
        // Remove the least recently used node from the list and map
        delete(lru.Cache, lru.Tail.Key)
        lru.removeFromList(lru.Tail)
    }
}
*/
//   Your LRUCache object will be instantiated and called as such:
func main() {
	capacity := 5
	key := 4
	value := 10
	obj := Constructor(capacity)
	param_1 := obj.Get(key)
	obj.Put(key, value)
	fmt.Println("param_1: ", param_1)
	param_2 := obj.Get(key)
	fmt.Println("param_2: ", param_2)

}
