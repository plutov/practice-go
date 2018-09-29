package brokennode

func FindBrokenNodes(brokenNodes int, reports []bool) string {

	result := make([]byte, len(reports))
	conf := make([]byte, len(reports))
	for i := range reports {
		conf[i] = 'B'
	}

out:
	for {
		for !check(conf, brokenNodes, reports) {
			if !next(conf) {
				break out
			}
		}

		merge(result, conf)
		if !next(conf) {
			break
		}
	}

	return string(result)
}

func merge(a, b []byte) {
	for i := range a {
		switch {
		case a[i] == 0:
			a[i] = b[i]
		case a[i] != b[i]:
			a[i] = '?'
		}
	}
}

func check(conf []byte, brokenNodes int, reports []bool) bool {

	// broken count
	broken := 0
	for i := range conf {
		if conf[i] == 'B' {
			broken++
		}
	}
	if broken != brokenNodes {
		return false
	}

	// consistence with reports
	n := len(conf)
	for i := range conf {
		if conf[i] == 'B' {
			continue
		}
		var want byte = 'B'
		if reports[i] {
			want = 'W'
		}
		if conf[(i+1)%n] != want {
			return false
		}
	}
	return true
}

func next(conf []byte) bool {
	for i := 0; i < len(conf); i++ {
		if conf[i] == 'B' {
			conf[i] = 'W'
			return true
		}
		conf[i] = 'B'
	}
	return false
}
