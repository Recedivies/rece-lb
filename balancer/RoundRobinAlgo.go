package balancer

type RoundRobinAlgo struct{}

var idx int = 0

func (rr RoundRobinAlgo) distribute(nodePool *NodePool) *Node {
	nodes := nodePool.np
	maxLen := len(nodes)

	nextNode := nodes[idx%maxLen]
	idx++

	return nextNode
}
