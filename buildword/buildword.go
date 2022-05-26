package buildword

import "strings"

func Search(adj [][]int, s int, t int, c int, res *int) {
	// TODO: bfs
	if s == t {
		if c < *res || *res == 0 {
			*res = c
		}
	} else {
		for _, v := range adj[s] {
			Search(adj, v, t, c+1, res)
		}
	}
}

func BuildWord(word string, fragments []string) int {
	adj := make([][]int, len(word))
	// TODO: trie
	for _, fragment := range fragments {
		i, j := 0, 0
		for j != -1 {
			j = strings.Index(word[i:], fragment)
			if j != -1 {
				i += j
				adj[i] = append(adj[i], i+len(fragment))
				i += len(fragment)
			}
		}
	}
	var res int
	Search(adj, 0, len(word), 0, &res)
	return res
}
