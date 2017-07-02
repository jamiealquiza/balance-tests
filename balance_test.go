package balancetests

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"testing"
)

const Nodes = 4

var (
	filePath string = "./words.txt"
	nodeList        = make([]string, Nodes)
)

func init() {
	for n := range nodeList {
		nodeList[n] = fmt.Sprintf("node-%d", n)
	}

	chInit()
	fnv1modInit()
	vaporchInit()
}

type method struct {
	name string
	f    func(string) string
}

var methods []method

// Words via https://raw.githubusercontent.com/dwyl/english-words/master/words.txt.

func TestBalance(t *testing.T) {
	for _, m := range methods {
		nodes := map[string]uint64{}

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			nodes[m.f(scanner.Text())]++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var total float64
		var counts []float64
		var empty int

		for _, v := range nodes {
			if v == 0 {
				empty++
			}
			vf := float64(v)
			total += vf
			counts = append(counts, vf)
		}

		sort.Float64s(counts)
		rng := counts[len(counts)-1] - counts[0]
		imbp := rng / total * 100
		imbr := counts[len(counts)-1] / counts[0]

		fmt.Printf("\n[%s]\n", m.name)
		fmt.Printf("%20s - %d\n", "Empty nodes", empty)
		fmt.Printf("%20s - portion of keys: %.2f%% / ratio: %.2fx\n",
			"Greatest imbalance", imbp, imbr)
		fmt.Printf("%20s - %.0f / highest value: %.0f / lowest value: %.0f\n",
			"Range", rng, counts[len(counts)-1], counts[0])

		if len(nodes) < 16 {
			fmt.Println()
			for _, n := range nodeList {
				fmt.Printf("%s: %d\n", n, nodes[n])
			}
		}
	}
}
