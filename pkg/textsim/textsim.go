package textsim

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

var tokenRe = regexp.MustCompile(`[A-Za-z0-9]+`)

func tokenize(s string) []string {
	s = strings.ToLower(s)
	toks := tokenRe.FindAllString(s, -1)
	return toks
}

func BuildTFIDF(docs []string) (vocab []string, vecs []map[string]float64, idf map[string]float64) {
	n := len(docs)
	if n == 0 {
		return nil, nil, nil
	}

	tf := make([]map[string]float64, n)
	df := make(map[string]int)
	for i, d := range docs {
		toks := tokenize(d)
		tf[i] = make(map[string]float64)
		for _, t := range toks {
			tf[i][t] += 1.0
		}

		seen := make(map[string]bool)
		for _, t := range toks {
			if !seen[t] {
				seen[t] = true
				df[t]++
			}
		}
	}

	for t := range df {
		vocab = append(vocab, t)
	}
	sort.Strings(vocab)

	idf = make(map[string]float64, len(df))
	for t, c := range df {
		idf[t] = math.Log(float64(n)/float64(1+c)) + 1.0
	}

	vecs = make([]map[string]float64, n)
	for i := 0; i < n; i++ {
		vecs[i] = make(map[string]float64)
		for term, cnt := range tf[i] {
			tfScaled := 1.0 + math.Log(cnt)
			vecs[i][term] = tfScaled * idf[term]
		}
	}

	for i := 0; i < n; i++ {
		norm := 0.0
		for _, val := range vecs[i] {
			norm += val * val
		}
		norm = math.Sqrt(norm)
		if norm > 0 {
			for k, v := range vecs[i] {
				vecs[i][k] = v / norm
			}
		}
	}

	return vocab, vecs, idf
}

func Cosine(a, b map[string]float64) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}
	dot := 0.0
	for k, va := range a {
		if vb, ok := b[k]; ok {
			dot += va * vb
		}
	}
	return dot
}

func PairwiseSimilarity(vecs []map[string]float64) [][]float64 {
	n := len(vecs)
	out := make([][]float64, n)
	for i := 0; i < n; i++ {
		out[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				out[i][j] = 1.0
			} else {
				out[i][j] = Cosine(vecs[i], vecs[j])
			}
		}
	}
	return out
}
