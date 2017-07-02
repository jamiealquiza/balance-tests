package balancetests

import (
	"github.com/jamiealquiza/ch"
)

var ring *ch.Ring

func chInit() {
	ring, _ = ch.New(&ch.Config{VNodes: 32})

	for _, n := range nodeList {
		ring.AddNode(n)
	}

	methods = append(methods, method{
		name: "consistent-hashing", f: ringGet})
}

func ringGet(k string) string {
	n, _ := ring.GetNode(k)
	return n
}
