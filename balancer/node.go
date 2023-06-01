package balancer

type Node struct {
	host string
}

// NewNode creates a new node
func NewNode(host string) *Node {
	return &Node{
		host: host,
	}
}
