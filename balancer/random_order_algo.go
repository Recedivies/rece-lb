package balancer

import (
	"errors"
	"math/rand"
)

type RandomOrderAlgo struct{}

func (rr RandomOrderAlgo) distribute(nodePool *NodePool) (*Node, error) {
	nodes := nodePool.np
	maxLen := len(nodes)

	randomNode := nodes[rand.Intn(maxLen)]
	if randomNode.getHealth() {
		return randomNode, nil
	}

	return nil, errors.New("service temporarily unavailable")
}
