package main

type Node struct {
	key, value int
	prev, next *Node
}



type LRUCache struct {
    cap int
	cache map[int]*Node
	head, tail *Node
}


func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	head.prev = head
    return LRUCache{
		cap : capacity, 
		cache:  make(map[int]* Node),
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) addNode (node *Node) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode (node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}


func (this *LRUCache) MoveToHead (node *Node) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
		this.MoveToHead(node)
		return node.value
	}
	return  -1
}

func (this *LRUCache) popTail () *Node{
	tail := this.tail.prev
	this.removeNode(tail)
	return tail 
}

func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.cache[key]; ok {
		node.value = value
		this.MoveToHead(node)
	}else {
		newNode := &Node {
			key: key,
			value: value,
		}
		this.cache[key] = newNode
		this.addNode(newNode)

		if len(this.cache) > this.cap{
			tail := this.popTail()
			delete(this.cache, tail.key)
		}
	}

}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */