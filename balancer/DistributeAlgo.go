package balancer

// define interface for strategy classes
type DistributeAlgo interface {
	distribute(nodePool *NodePool) *Node
}
