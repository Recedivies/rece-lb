package balancer

type LeastConnectionAlgo struct{}

func (rr LeastConnectionAlgo) distribute(nodePool *NodePool) *Node {
	var leastConnectionNode *Node
	var numOfConnection uint32 = 1<<32 - 1

	for _, node := range nodePool.np {
		if numOfConnection > node.requestCount {
			numOfConnection = node.requestCount
			leastConnectionNode = node
		}
	}

	return leastConnectionNode
}
