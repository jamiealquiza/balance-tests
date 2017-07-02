package balancetests

import (
	"strconv"

	"github.com/jamiealquiza/ch"
)

var ring *ch.Ring

func init() {
	ring, _ = ch.New(&ch.Config{VNodes: 1})

	for i := 0; i < Nodes; i++ {
		ring.AddNode(strconv.Itoa(i))
	}

	methods = append(methods, method{
		name: "ch", f: ringGet})
}

func ringGet(k string) int {
	n, _ := ring.GetNode(k)
	i, _ := strconv.Atoi(n)
	return i
}
