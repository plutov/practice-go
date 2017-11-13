package missingnumbers

// My solution
//
// Store each value as a bit in an array
// of 64 bit "words". Then find all words which
// hasn't all bits set. And in that word find the missing bits.
// So bit 66 is stored in word 1 bit 2

// How big is a word (64 since I use uint64)
const wordSize = 64
// a 64-bit number with all bits set
const wordAllSet = 0xffffffffffffffff
// Mask to use instead of modulo wordSize
const moduloMask = 63

type bitArray struct {
	len int // Total number of bits
	nWords int // THe number of words used
	words []uint64
}

func newBitArray(len int) bitArray {
	nWords := len / wordSize
	if len & moduloMask != 0 {
		nWords++
	}
	words := make([]uint64, nWords, nWords)
	return bitArray{len: len, nWords: nWords, words: words}
}

func (b *bitArray) setBit(n int) {
	index := n / wordSize
	bit := uint64(1) << uint(n & moduloMask)
	b.words[index] |= bit
}

func (b *bitArray) findUnsetBits(maxCount int) []int {
	unset := make([]int, 0, maxCount)
	for i, word := range b.words {
		if word & wordAllSet != wordAllSet {
			j := 0
			for j < wordSize {
				if word & (uint64(1) << uint(j)) == 0 {
					unset = append(unset, i * wordSize + j)
					if len(unset) == maxCount {
						return unset
					}
				}
				j++
			}
		}
	}
	return unset
}

func Missing(numbers []int) []int {
	bits := newBitArray(len(numbers) + 3) // 0 and 2 missing numbers
	bits.setBit(0)
	for _, number := range numbers {
		bits.setBit(number)
	}
	return bits.findUnsetBits(2)
}

