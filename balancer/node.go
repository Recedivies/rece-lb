package balancer

import (
	"sync"
)

type Node struct {
	Host         string
	RequestCount uint32
	Health       bool
}

var mutex = &sync.RWMutex{}

// NewNode creates a new node
func NewNode(host string) *Node {
	return &Node{
		Host:         host,
		RequestCount: 0,
		Health:       true,
	}
}

// getHealth returns the value of Health in Node.
func (node *Node) getHealth() bool {
	mutex.RLock()
	health := node.Health
	mutex.RUnlock()
	return health
}
