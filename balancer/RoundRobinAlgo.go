package balancer

import "sync"

type RoundRobinAlgo struct{}

var idx int = 0
var mutex = sync.RWMutex{}

func (rr RoundRobinAlgo) distribute(nodePool *NodePool) *Node {
	// use mutex to prevent data race condition
	mutex.Lock()

	nodes := nodePool.np
	maxLen := len(nodes)

	nextNode := nodes[idx%maxLen]
	idx++

	mutex.Unlock()

	return nextNode
}
