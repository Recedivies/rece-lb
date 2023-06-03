package balancer

import (
	"log"
	"net"
	"net/url"
	"sync"
	"time"
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

// ping checks if the node is alive.
func (node *Node) isAlive() bool {
	u, _ := url.Parse(node.Host)
	conn, err := net.DialTimeout("tcp", u.Host, 1*time.Second)
	if err != nil {
		log.Println("node is unreachable: ", err)
		return false
	}
	defer conn.Close()
	return true
}

// markStatus marks the status of Health in Node.
func (node *Node) markStatus(status bool) {
	mutex.Lock()
	node.Health = status
	mutex.Unlock()
}

// getHealth returns the value of Health in Node.
func (node *Node) getHealth() bool {
	mutex.RLock()
	health := node.Health
	mutex.RUnlock()
	return health
}
