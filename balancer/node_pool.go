package balancer

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
)

const ErrBackendUnreachable = "unsupported protocol"

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
		nextNode, err := nodePool.distributionAlgo.distribute(nodePool)
		if err != nil {
			return &Node{}, err
		}

		return nextNode, nil
	}

	if nodePool.algorithm == "Random" {
		nodePool.setDistributionAlgo(RandomOrderAlgo{})
		node, err := nodePool.distributionAlgo.distribute(nodePool)
		if err != nil {
			return &Node{}, err
		}

		return node, nil
	}

	if nodePool.algorithm == "LeastConnection" {
		nodePool.setDistributionAlgo(LeastConnectionAlgo{})
		leastConnectionNode, err := nodePool.distributionAlgo.distribute(nodePool)
		if err != nil {
			return &Node{}, err
		}

		return leastConnectionNode, nil
	}

	return &Node{}, errors.New("algorithm not supported")
}

func (nodePool *NodePool) director(req *http.Request) {
	node, err := nodePool.GetNode()
	// Check backend server health
	if err != nil {
		req.URL = &url.URL{}
		return
	}

	// increment the number of request the server is serving
	atomic.AddUint32(&node.RequestCount, 1)

	u, _ := url.Parse(node.Host)

	req.URL.Scheme = u.Scheme
	req.URL.Host = u.Host
}

func (nodePool *NodePool) modifyResponse(res *http.Response) error {
	fmt.Println(res.StatusCode)
	for _, n := range nodePool.np {
		u, _ := url.Parse(n.Host)
		if u.Host == res.Request.URL.Host {
			// decrement the number of request the host is serving
			atomic.AddUint32(&n.RequestCount, ^uint32(n.RequestCount-1))
			log.Println(n)
			break
		}
	}
	return nil
}

func (nodePool *NodePool) errorHandler(w http.ResponseWriter, req *http.Request, err error) {
	if strings.Contains(err.Error(), ErrBackendUnreachable) {
		http.Error(w, "Backend server is currently unreachable", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}
