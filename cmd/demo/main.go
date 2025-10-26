package main

import (
	"fmt"

	"github.com/majorbruteforce/hifive/pkg/textsim"
)

func main() {
	docs := []string{
		"I like stars.",
		"Can we meet today?.",
		"I went to the store today.",
		"Stars are awesome!",
	}

	_, vecs, _ := textsim.BuildTFIDF(docs)
	sim := textsim.PairwiseSimilarity(vecs)

	for i := 0; i < len(docs); i++ {
		for j := i + 1; j < len(docs); j++ {
			fmt.Printf("%d-%d: %.4f\n", i, j, sim[i][j])
		}
	}
}
