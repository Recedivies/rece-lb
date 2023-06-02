package balancer

import (
	"errors"
)

type LeastConnectionAlgo struct{}

func (rr LeastConnectionAlgo) distribute(nodePool *NodePool) (*Node, error) {
	var leastConnectionNode *Node
	nodes := nodePool.np
	var numOfConnection uint32 = 1<<32 - 1

	for _, node := range nodes {
		if numOfConnection > node.RequestCount && node.getHealth() {
			numOfConnection = node.RequestCount
			leastConnectionNode = node
		}
	}

	if leastConnectionNode == nil {
		return nil, errors.New("service temporarily unavailable")
	}

	return leastConnectionNode, nil
}
