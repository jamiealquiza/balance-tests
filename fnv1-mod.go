package balancetests

func init() {
	methods = append(methods,
		method{name: "FNV1-mod", f: fnv1mod},
		method{name: "FNV1a-mod", f: fnv1amod})
}

func fnv1mod(k string) int {
	var h = 2166136261
	for _, c := range []byte(k) {
		h *= 16777619
		h ^= int(c)
	}

	return h & (Nodes - 1)
}

func fnv1amod(k string) int {
	var h = 2166136261
	for _, c := range []byte(k) {
		h ^= int(c)
		h *= 16777619
	}

	return h & (Nodes - 1)
}
