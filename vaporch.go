package balancetests

import (
	"github.com/jamiealquiza/vaporch"
)

var vch *vaporch.Ring

func vaporchInit() {
	vch, _ = vaporch.New(&vaporch.Config{
		Nodes: nodeList,
	})

	methods = append(methods, method{
		name: "vaporCH", f: vchGet})
}

func vchGet(k string) string {
	return vch.Get(k)
}
