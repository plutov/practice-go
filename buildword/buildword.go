package buildword

import (
	"github.com/eapache/queue"
	"strings"
)

type Pair struct {
	s int
	d int
}

func bfs(adj [][]int, s int, t int) int {
	vis := make([]bool, len(adj))
	vis[s] = true
	q := queue.New()
	q.Add(Pair{s, 0})
	for q.Length() > 0 {
		p := q.Peek().(Pair)
		s = p.s
		d := p.d
		if s == t {
			return d
		}
		q.Remove()
		vis[s] = true
		for _, v := range adj[s] {
			if !vis[v] {
				q.Add(Pair{v, d + 1})
			}
		}
	}
	return 0
}

func BuildWord(word string, fragments []string) int {
	adj := make([][]int, len(word)+1)
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
	return bfs(adj, 0, len(word))
}
