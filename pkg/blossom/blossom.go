package blossom

/*
#cgo CXXFLAGS: -std=c++11 -O3
#cgo LDFLAGS: -L${SRCDIR}/../../blossom5-v2.05.src/lib -lblossom
#cgo LDFLAGS: -L${SRCDIR} -lblossom_wrapper
#cgo LDFLAGS: -lstdc++

#include "blossom_wrapper.hpp"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// Handle is a Go wrapper around the Blossom C handle
type Handle struct {
	ptr unsafe.Pointer
}

// New creates a new Blossom solver instance
func New(numNodes, numEdges int) *Handle {
	h := &Handle{ptr: C.blossom_new(C.int(numNodes), C.int(numEdges))}
	runtime.SetFinalizer(h, func(h *Handle) { h.Free() })
	return h
}

// AddEdge adds an edge between u and v with weight w
func (h *Handle) AddEdge(u, v int, w float64) {
	C.blossom_add_edge(h.ptr, C.int(u), C.int(v), C.double(w))
}

// Solve runs the Blossom algorithm
func (h *Handle) Solve() {
	C.blossom_solve(h.ptr)
}

// GetMatch returns the node matched with i
func (h *Handle) GetMatch(i int) int {
	return int(C.blossom_get_match(h.ptr, C.int(i)))
}

// Free releases the native memory
func (h *Handle) Free() {
	if h.ptr != nil {
		C.blossom_free(h.ptr)
		h.ptr = nil
	}
}
