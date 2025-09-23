#pragma once

#ifdef __cplusplus
extern "C" {
#endif

void* blossom_new(int num_nodes, int num_edges);

void blossom_add_edge(void* handle, int u, int v, double weight);

void blossom_solve(void* handle);

int blossom_get_match(void* handle, int i);

void blossom_free(void* handle);

#ifdef __cplusplus
}
#endif
