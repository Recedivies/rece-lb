package balancer

import (
	"net/http/httputil"
)

// load balancer implemented on top of built in httpUtil.ReverseProxy
type LoadBalancer struct {
	*httputil.ReverseProxy
	np *NodePool
}

// NewLoadBalancer creates a new node pool
func NewLoadBalancer(servers []string, algorithm string) *LoadBalancer {
	np := NewNodePool(servers, algorithm)

	lb := &LoadBalancer{
		np: np,
		ReverseProxy: &httputil.ReverseProxy{
			// this function is called before sending request to nodes
			Director:       np.director,
			ModifyResponse: np.modifyResponse,
			ErrorHandler:   np.errorHandler,
		},
	}

	return lb
}
