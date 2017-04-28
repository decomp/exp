package x86

import "github.com/decomp/exp/bin"

// queue represents a queue of addresses.
type queue struct {
	// Addresses in the queue.
	addrs map[bin.Address]bool
}

// newQueue returns a new queue.
func newQueue() *queue {
	return &queue{
		addrs: make(map[bin.Address]bool),
	}
}

// push pushes the given address to the queue.
func (q *queue) push(addr bin.Address) {
	q.addrs[addr] = true
}

// pop pops an address from the queue.
func (q *queue) pop() bin.Address {
	if len(q.addrs) == 0 {
		panic("invalid call to pop; empty queue")
	}
	var min bin.Address
	for addr := range q.addrs {
		if min == 0 || addr < min {
			min = addr
		}
	}
	delete(q.addrs, min)
	return min
}

// empty reports whether the queue is empty.
func (q *queue) empty() bool {
	return len(q.addrs) == 0
}
