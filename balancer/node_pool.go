package balancer

import (
	"errors"
	"log"
	"net/http"
	"net/url"
)

// A pool of all hosts and other conf
type NodePool struct {
	np               []*Node
	algorithm        string
	distributionAlgo DistributionAlgo
}

func (nodePool *NodePool) setDistributionAlgo(d DistributionAlgo) {
	nodePool.distributionAlgo = d
}

// NewNodePool creates a new node pool
func NewNodePool(servers []string, algorithm string) *NodePool {
	var nodes []*Node
	for _, server := range servers {
		nodes = append(nodes, NewNode(server))
	}

	return &NodePool{
		np:        nodes,
		algorithm: algorithm,
	}
}

// GetNode gets a new Node available to forward request to based on algorithm
func (nodePool *NodePool) GetNode() (*Node, error) {
	if nodePool.algorithm == "RoundRobin" {
		nodePool.setDistributionAlgo(RoundRobinAlgo{})
		nextNode := nodePool.distributionAlgo.distribute(nodePool)

		return nextNode, nil
	}

	if nodePool.algorithm == "Random" {
		nodePool.setDistributionAlgo(RandomOrderAlgo{})
		node := nodePool.distributionAlgo.distribute(nodePool)

		return node, nil
	}

	return &Node{}, errors.New("algorithm not supported")
}

func (nodePool *NodePool) Director(req *http.Request) {
	node, _ := nodePool.GetNode()
	log.Println(node)

	u, _ := url.Parse(node.host)

	req.URL.Scheme = u.Scheme
	req.URL.Host = u.Host
}
