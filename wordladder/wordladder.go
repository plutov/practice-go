package wordladder

import (
	"github.com/eapache/queue"
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
			return d + 1
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

func dist(str1 string, str2 string) bool {
	res := false
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			if res {
				return false
			}
			res = true
		}
	}
	return res
}

func add(str string, dic *[]string) int {
	for i, s := range *dic {
		if s == str {
			return i
		}
	}
	*dic = append(*dic, str)
	return len(*dic) - 1
}

func WordLadder(from string, to string, dic []string) int {
	s := add(from, &dic)
	t := add(to, &dic)
	adj := make([][]int, len(dic))
	for i := 0; i < len(dic); i++ {
		for j := i + 1; j < len(dic); j++ {
			if dist(dic[i], dic[j]) {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}
	return bfs(adj, s, t)
}
