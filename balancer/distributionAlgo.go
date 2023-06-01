package balancer

// define interface for strategy classes
type DistributionAlgo interface {
	distribute(nodePool *NodePool) *Node
}
