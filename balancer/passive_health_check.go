package balancer

import (
	"log"
	"time"
)

func (nodePool *NodePool) passiveHeathCheck() {
	for _, node := range nodePool.np {
		log.Println("Health check for host :", node.Host)
		go node.markStatus(node.isAlive())
	}
}

// StartPassiveHeathCheck starts passive health check specified interval
func (lb *LoadBalancer) StartPassiveHeathCheck() {
	t := time.NewTicker(time.Second * time.Duration(lb.np.healthCheckInterval))
	for {
		select {
		case <-t.C:
			log.Println("Starting passive health check...")
			lb.np.passiveHeathCheck()
			log.Println("Passive Health check completed")
		}
	}
}
