package alien_dictionary

import "fmt"

func debugOutEdges(edges map[uint8]map[uint8]bool) {
	fmt.Printf("=== DEBUG IN OUTEDGES ===\n")
	for k, v := range edges {
		fmt.Printf("%s: [", string(k))
		first := true
		for to := range v {
			if !first {
				fmt.Printf(" ")
			}
			fmt.Printf("%s", string(to))
			first = false
		}
		fmt.Printf("]\n")
	}
}

func debugInDegree(degree map[uint8]int) {
	fmt.Printf("=== DEBUG IN DEGREE ===\n")
	for k, v := range degree {
		fmt.Printf("%s: %v\n", string(k), v)
	}
}
