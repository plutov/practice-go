package lastlettergame

func Sequence(dic []string) []string {
	p := newPathFinder(dic)
	result := p.findLongest()

	return result
}

func newPathFinder(dic []string) *pathFinder {
	return &pathFinder{
		dic:     dic,
		lookup:  buildLookup(dic),
		visited: make([]bool, len(dic)+1),
		result:  make([]string, 0, len(dic)),
	}
}

type pathFinder struct {
	dic     []string
	result  []string
	lookup  [][]int
	visited []bool
}

func buildLookup(dic []string) [][]int {
	l := len(dic)
	lookup := make([][]int, l+1)

	for io, wo := range dic {
		if len(lookup[io]) == 0 {
			lookup[io] = make([]int, 0, l)
		}

		for ii, wi := range dic {
			if wo[len(wo)-1] == wi[0] && wo != wi {
				lookup[io] = append(lookup[io], ii)
			}
		}
	}

	lookup[l] = make([]int, l)
	for i := 0; i < l; i++ {
		lookup[l][i] = i
	}

	return lookup
}

func (p *pathFinder) enter(i int) {
	p.visited[i] = true
}

func (p *pathFinder) exit(index int) {
	p.visited[index] = false
}

func (p *pathFinder) isVisited(i int) bool {
	return p.visited[i]
}

func (p *pathFinder) setResult(result []string) {
	p.result = p.result[:len(result)]
	copy(p.result, result)
}

func (p *pathFinder) find(currentIndex int, rest []string) []string {
	for _, nextIndex := range p.lookup[currentIndex] {
		if p.isVisited(nextIndex) {
			continue
		}

		p.enter(nextIndex)

		candidate := append(rest, p.dic[nextIndex])
		candidate = p.find(nextIndex, candidate)

		p.exit(nextIndex)

		if len(candidate) > len(p.result) {
			p.setResult(candidate)
		}
	}

	return rest
}

func (p *pathFinder) findLongest() []string {
	p.find(len(p.dic), make([]string, 0, len(p.dic)))
	return p.result
}
