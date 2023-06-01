package balancer

type Node struct {
	host         string
	requestCount uint32
}

// NewNode creates a new node
func NewNode(host string) *Node {
	return &Node{
		host: host,
	}
}
