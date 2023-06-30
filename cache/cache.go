// package cache
package main

import "fmt"

const SIZE int = 5

// HASH this can be used by any type
type HASH map[int]*Node

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	HASH  HASH
}

// NewQueue create a empty queue
func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	head.Left = head
	return Queue{Head: head, Tail: tail}
}

// New crete a new empty cache
func New() Cache {
	return Cache{Queue: NewQueue(), HASH: HASH{}}
}

// Check if value already exist in cache. If y, so remove from cache and put it in the cache again.
// But, if not already exist in the cache, so it's will be added
func (c *Cache) Check(k int) {
	node := &Node{}

	if value, exist := c.HASH[k]; exist {
		node = c.Remove(value)
	} else {
		node = &Node{Value: k}
	}
	c.Add(node)
	c.HASH[k] = node
}

// Add value in the cache
func (c *Cache) Add(node *Node) {
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = node
	node.Left = c.Queue.Head
	node.Right = tmp
	tmp.Left = node

	c.Queue.Length++

	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

// Remove value from the cache, returning a Node type
func (c *Cache) Remove(value *Node) *Node {
	left := value.Left
	right := value.Right

	left.Right = right
	right.Left = left
	c.Queue.Length--

	delete(c.HASH, value.Value)
	return value
}

func (c Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	aNode := q.Head.Right
	for i := 0; i < q.Length; i++ {
		fmt.Printf("%v \n", aNode.Value)
		aNode = aNode.Right
	}
}

func main() {
	values := []int{19, 14}
	cache := New()

	for _, value := range values {
		cache.Check(value)
		cache.Display()
	}
}
