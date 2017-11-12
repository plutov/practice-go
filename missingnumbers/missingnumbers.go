package missingnumbers

// My solution
//
// Store each value as a bit in an array
// of 64 bit "words". Then find all words which
// hasn't all bits set. And in that word find the missing bits.
// I cheat a bit because the last word should be treated differently
// But it works now then I know two bits are unset.

// How big is a word (64 since I use uint64)
const wordSize = 64
// a 64-bit number with all bits set
const wordAllSet = 0xffffffffffffffff

// An array of all single bit set numbers
var wordBits = calcWordBits()

// Calculate an array of bits 1, 2, 4 and so on
// Used to and with word to see if a bit is set
func calcWordBits() [wordSize]uint64 {
	wordBits := [wordSize]uint64{}
	i := uint(0)
	for i < wordSize {
		wordBits[i] = 1 << i
		i++
	}
	return wordBits
}


type bitArray struct {
	len int // Total number of bits
	nWords int // THe number of words used
	words []uint64

}

func newBitArray(len int) bitArray {
	nWords := len / wordSize
	// @TODO use mask instead
	if len % wordSize != 0 {
		nWords++
	}
	words := make([]uint64, nWords, nWords)
	return bitArray{len: len, nWords: nWords, words: words}
}

func (b *bitArray) setBit(n int) {
	index := n / wordSize
	// @TODO use mask instead
	bit := uint64(1) << uint(n % wordSize)
	b.words[index] |= bit
}


func (b *bitArray) findUnsetBits(maxCount int) []int {
	unset := make([]int, 0, maxCount)
	for i, word := range b.words {
		if word & wordAllSet != wordAllSet {
			for j, wordBit := range wordBits {
				if word & wordBit == 0 {
					unset = append(unset, i * wordSize + j)
					if len(unset) == maxCount {
						return unset
					}
				}
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
