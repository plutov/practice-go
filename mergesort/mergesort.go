package mergesort

// MergeSort sorts a slice of ints
func MergeSort(input []int) []int {
    if len(input) <= 1 {
        return input
    }

    return merge(
        MergeSort(input[:len(input)/2]),
        MergeSort(input[len(input)/2:]),
    )
}

func merge(l []int, r []int) []int {
    var idx, idxL, idxR int

    result := make([]int, len(l)+len(r))

    for {
        if l[idxL] <= r[idxR] {
            result[idx] = l[idxL]
            idxL++
        } else {
            result[idx] = r[idxR]
            idxR++
        }

        idx++

        // left slice empty
        if len(l) == idxL {
            copy(result[idx:], r[idxR:])
            break
        }
        // right slice empty
        if len(r) == idxR {
            copy(result[idx:], l[idxL:])
            break
        }
    }

    return result
}
