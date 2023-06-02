package balancer

import "errors"

type RoundRobinAlgo struct{}

var currentServerIndex int = 0

func (rr RoundRobinAlgo) distribute(nodePool *NodePool) (*Node, error) {
	// use mutex to prevent data race condition
	mutex.Lock()

	var nextNode *Node
	nodes := nodePool.np
	maxLen := len(nodes)

	for i := currentServerIndex; i <= currentServerIndex+maxLen; i++ {
		currentIdx := i % maxLen
		nextNode = nodes[currentIdx]
		currentServerIndex++

		if nextNode.getHealth() {
			return nextNode, nil
		}
	}

	mutex.Unlock()

	return nil, errors.New("service temporarily unavailable")
}
