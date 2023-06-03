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
func NewLoadBalancer(servers []string, algorithm, healthCheckType string, healthCheckInterval int) *LoadBalancer {
	np := NewNodePool(servers, algorithm, healthCheckType, healthCheckInterval)

	lb := &LoadBalancer{
		np: np,
		ReverseProxy: &httputil.ReverseProxy{
			// this function is called before sending request to nodes
			Director:       np.director,
			ModifyResponse: np.modifyResponse,
			ErrorHandler:   np.errorHandler,
		},
	}
	if healthCheckType == "passive" {
		// starts passive health check if opted
		go lb.StartPassiveHeathCheck()
	}

	return lb
}
