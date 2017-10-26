package secretmessage

import (
	"fmt"
	"strconv"
	"strings"
)

func InsertionSort(arg []string) {
	i := 1
	j := 0
	var x string
	var y, xval int
	for i < len(arg) {
		x = arg[i]
		j = i - 1
		y, _ = strconv.Atoi(strings.Split(arg[j], ":")[1])
		xval, _ = strconv.Atoi(strings.Split(x, ":")[1])
		for j >= 0 && y > xval {
			arg[j+1] = arg[j]
			j -= 1
			if j >= 0 {
				y, _ = strconv.Atoi(strings.Split(arg[j], ":")[1])
			}
		}
		arg[j+1] = x
		i += 1
	}
}

func Encode(encoded string) string {
	var keys []string
	var counter map[string]int
	var currentKey string

	counter = make(map[string]int)

	for _, runeVal := range encoded {
		currentKey = string(runeVal)
		_, inMap := counter[currentKey]
		if !inMap {
			counter[currentKey] = 1
			keys = append(keys, currentKey)
		} else {
			counter[currentKey] = counter[currentKey] + 1
		}
	}

	var temp []string
	for k, v := range counter {
		temp = append(temp, k+":"+fmt.Sprintf("%d", v))
	}
	InsertionSort(temp)

	var g string
	var result []string

	for z := len(temp) - 1; z >= 0; z-- {
		g = strings.Split(temp[z], ":")[0]
		if g == "_" {
			break
		} else {
			result = append(result, g)
		}
	}

	return strings.Join(result, "")
}
