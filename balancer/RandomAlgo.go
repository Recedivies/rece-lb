package balancer

import "math/rand"

type RandomAlgo struct{}

func (rr RandomAlgo) distribute(nodePool *NodePool) *Node {

	nodes := nodePool.np
	maxLen := len(nodes)

	nextNode := nodes[rand.Intn(maxLen)]

	return nextNode
}
