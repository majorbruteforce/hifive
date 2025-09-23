package main

import (
	"fmt"

	"github.com/majorbruteforce/hifive/pkg/blossom"
)

func main() {
	h := blossom.New(4, 4)
	defer h.Free()

	h.AddEdge(0, 1, 1.0)
	h.AddEdge(1, 2, 1.0)
	h.AddEdge(2, 3, 1.0)
	h.AddEdge(3, 0, 1.0)

	h.Solve()

	for i := 0; i < 4; i++ {
		fmt.Printf("node %d matched with %d\n", i, h.GetMatch(i))
	}
}
