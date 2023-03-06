package algo

import "fmt"

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func test146() {
	obj := Constructor(2)
	//obj.Check()
	//fmt.Println("obj.Put(1, 1)")
	obj.Put(1, 1)
	//obj.Check()
	//fmt.Println("obj.Put(2, 2)")
	obj.Put(2, 2)
	//obj.Check()
	fmt.Println("obj.Get(1)", obj.Get(1))
	//obj.Check()
	//fmt.Println("obj.Put(3, 3)")
	obj.Put(3, 3)
	//obj.Check()
	fmt.Println("obj.Get(2)", obj.Get(2))
	//obj.Check()
	obj.Put(4, 4)
	//obj.Check()
	fmt.Println("obj.Get(1)", obj.Get(1))
	//obj.Check()
	fmt.Println("obj.Get(3)", obj.Get(3))
	//obj.Check()
	fmt.Println("obj.Get(4)", obj.Get(4))
	//obj.Check()
}

type LruListNode struct {
	Prev *LruListNode
	Next *LruListNode
	key  int
	Val  int
}

type LRUCache struct {
	capacity int
	length   int
	head     *LruListNode
	tail     *LruListNode
	hashMap  map[int]*LruListNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &LruListNode{}, &LruListNode{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		capacity: capacity,
		length:   0,
		head:     head,
		tail:     tail,
		hashMap:  make(map[int]*LruListNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if this.length <= 0 {
		return -1
	}
	v, ok := this.hashMap[key]
	if !ok {
		return -1
	}
	v.Prev.Next = v.Next
	v.Next.Prev = v.Prev
	v.Next = this.head.Next
	v.Prev = this.head
	this.head.Next.Prev = v
	this.head.Next = v
	return v.Val
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.hashMap[key]
	if ok {
		v.Val = value
		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		v.Next = this.head.Next
		v.Prev = this.head
		this.head.Next.Prev = v
		this.head.Next = v
		return
	}
	node := &LruListNode{
		Prev: this.head,
		Next: this.head.Next,
		key:  key,
		Val:  value,
	}
	this.head.Next.Prev = node
	this.head.Next = node
	this.hashMap[key] = node
	this.length++
	if this.length > this.capacity {
		delete(this.hashMap, this.tail.Prev.key)
		this.tail.Prev.Prev.Next = this.tail
		this.tail.Prev = this.tail.Prev.Prev
		this.length--
	}

}

func (this *LRUCache) Check() {
	fmt.Printf("len:%d, node:[", this.length)
	for p := this.head; p != nil; p = p.Next {
		fmt.Printf(" %d ", p.Val)
	}
	fmt.Printf("], [")
	for p := this.tail; p != nil; p = p.Prev {
		fmt.Printf(" %d ", p.Val)
	}
	fmt.Printf("], hashMap:[")
	for k, _ := range this.hashMap {
		//fmt.Printf(" %d:%v, ", k, v)
		fmt.Printf(" %d ", k)
	}
	fmt.Printf("]\n")
}
