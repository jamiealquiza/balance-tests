package balancetests

import (
	"strconv"

	"github.com/jamiealquiza/vaporch"
)

var vch *vaporch.Ring

func init() {
	nodes := []string{}

	for i := 0; i < Nodes; i++ {
		nodes = append(nodes, strconv.Itoa(i))
	}

	vch, _ = vaporch.New(&vaporch.Config{
		Nodes: nodes,
	})

	methods = append(methods, method{
		name: "vaporCH", f: vchGet})
}

func vchGet(k string) int {
	i, _ := strconv.Atoi(vch.Get(k))
	return i
}
