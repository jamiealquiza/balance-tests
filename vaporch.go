package balancetests

import (
	"strconv"

	"github.com/jamiealquiza/vaporch"
)

var ch *vaporch.Ring

func init() {
	nodes := []string{}

	for i := 0; i < Nodes; i++ {
		nodes = append(nodes, strconv.Itoa(i))
	}

	ch, _ = vaporch.New(&vaporch.Config{
		VNodes: 3,
		Nodes:  nodes,
	})

	methods = append(methods, method{
		name: "vaporCH", f: vch})
}

func vch(k string) int {
	i, _ := strconv.Atoi(ch.Get(k))
	return i
}
