package balancer

import "math/rand"

type RandomOrderAlgo struct{}

func (rr RandomOrderAlgo) distribute(nodePool *NodePool) *Node {
	nodes := nodePool.np
	maxLen := len(nodes)

	nextNode := nodes[rand.Intn(maxLen)]

	return nextNode
}
