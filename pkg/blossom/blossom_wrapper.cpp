#include "blossom_wrapper.hpp"
#include "PerfectMatching.h"

#include <vector>

struct BlossomHandle {
    PerfectMatching* pm;
    int num_nodes;
    int num_edges;
    std::vector<int> edges;
    std::vector<double> weights;
};

extern "C" {

void* blossom_new(int num_nodes, int num_edges) {
    BlossomHandle* h = new BlossomHandle();
    h->num_nodes = num_nodes;
    h->num_edges = num_edges;
    h->pm = new PerfectMatching(num_nodes, num_edges);
    h->edges.reserve(2*num_edges);
    h->weights.reserve(num_edges);
    return h;
}

void blossom_add_edge(void* handle, int u, int v, double weight) {
    BlossomHandle* h = static_cast<BlossomHandle*>(handle);
    h->pm->AddEdge(u, v, weight);
}

void blossom_solve(void* handle) {
    BlossomHandle* h = static_cast<BlossomHandle*>(handle);
    h->pm->Solve();
}

int blossom_get_match(void* handle, int i) {
    BlossomHandle* h = static_cast<BlossomHandle*>(handle);
    return h->pm->GetMatch(i);
}

void blossom_free(void* handle) {
    BlossomHandle* h = static_cast<BlossomHandle*>(handle);
    delete h->pm;
    delete h;
}

} // extern "C"
