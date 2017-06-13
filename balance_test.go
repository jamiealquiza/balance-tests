package balancetests

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"testing"
)

const Nodes = 1024

var (
	filePath string = "./words.txt"
)

type method struct {
	name string
	f    func(string) int
}

var methods []method

// Words via https://raw.githubusercontent.com/dwyl/english-words/master/words.txt.

func TestBalance(t *testing.T) {
	for _, m := range methods {
		nodes := map[int]uint64{}

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
		for k := range nodes {
			if k == 0 {
				empty++
			}
			v := float64(nodes[k])
			total += v
			counts = append(counts, v)
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
	}
}
