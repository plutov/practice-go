package compression

import (
	"encoding/binary"
	"fmt"
	"math/bits"
)

func Encode(s string) string {
	// LZ77 compression, with no entropy coding.
	//
	// The idea is to find repeated sequences of bytes and replace them with
	// a reference to the previous occurrence.
	//
	// We use two hash tables, one for 8-byte sequences and one for 4-byte sequences.
	// This is a very simple and inefficient implementation, but it's easy to understand.
	//
	// The format is:
	// - A uvarint with the length of the next sequence.
	// - If the lowest bit is 0, then the next bytes are literals.
	// - If the lowest bit is 1, then the next bytes are a reference to a previous sequence.
	//   - The next uvarint is the offset of the match into previously decoded bytes from current position.
	//
	// For higher efficiency with the same encoding scheme,
	// some amount of "chaining" could be used, where more candidates are checked.
	//
	const printDebug = false

	// Since we only do ~2MB, make a lossless table.
	// This is very slow and memory consuming, but we do it for simplicity.
	// 4/8 byte value -> last offset
	var longTab = make(map[uint64]int, len(s))
	var shortTab = make(map[uint32]int, len(s)/8)

	// Source, src index
	src, srci := []byte(s), 0
	dst := make([]byte, 0, len(src))
	nextEmit := 0

	// Emit any pending literals
	emitLits := func() {
		if srci == nextEmit {
			return
		}
		if printDebug {
			fmt.Println("Emitting", srci-nextEmit, "literals.")
		}

		// Literals have the lowest bit set to 0
		l := uint64(srci-nextEmit) << 1
		dst = binary.AppendUvarint(dst, l)
		dst = append(dst, src[nextEmit:srci]...)
	}
	// Index bytes from 'from' (inclusive) to 'to' (exclusive)
	index := func(from, to int) {
		if to >= len(src)-8 {
			// Skip at end of block.
			return
		}
		for i := from; i < to; i++ {
			cur := binary.LittleEndian.Uint64(src[i : i+8])
			longTab[cur] = i
			shortTab[uint32(cur)] = i
		}
	}
	// Stop when we can't read 8 bytes
	for srci < len(src)-8 {
		if len(dst) > len(src)-8 {
			// Incompressible, just emit as literals
			dst = dst[:0]
			nextEmit = 0
			srci = len(src)
			emitLits()
			// We emitted everything, update the next emit
			nextEmit = srci
			break
		}
		cur := binary.LittleEndian.Uint64(src[srci : srci+8])

		// Read a long and short candidate.
		// Since we have a "perfect" hash table we know the destination matches.
		longOff, longOK := longTab[cur]
		shortOff, shortOK := shortTab[uint32(cur)]
		if longOK {
			// We match at least 8 bytes.
			// Fine enough for us to emit a match without checking more.
			// Emit any pending literals
			emitLits()

			// Find match length. We already know 8 bytes match.
			length := matchLen(src[srci+8:], src[longOff+8:]) + 8

			// Add length as uvarint
			// Matches have the lowest bit set to 1
			l := uint64(length)<<1 + 1
			dst = binary.AppendUvarint(dst, l)
			dst = binary.AppendUvarint(dst, uint64(srci-longOff))
			if printDebug {
				fmt.Println("Emitting long offset", srci-longOff, "length", length)
			}
			// Index added bytes
			index(srci, srci+length)
			srci += length
			nextEmit = srci
			continue
		}
		if shortOK {
			// Find match length. We already know 4 bytes match, and
			// we know we will not match 8 bytes, so as an optimization this could be
			//	length := bits.TrailingZeros64(cur ^ binary.LittleEndian.Uint64(src[shortOff:])) / 8
			length := matchLen(src[srci+4:], src[shortOff+4:]) + 4

			// Try finding a longer match at our position + 1.
			// This demonstrates very basic "lazy matching",
			// where we emit one literal if we can find a better match instead of the short one.
			if srci < len(src)-9 {
				cur := binary.LittleEndian.Uint64(src[srci+1 : srci+9])
				if longOff, longOK := longTab[cur]; longOK {
					lengthLong := matchLen(src[srci+9:], src[longOff+8:]) + 8
					// Emitting a literal will be the 1 (the literal itself) +
					// the length of the match, which may add one byte.
					if lengthLong > length+2 {
						if printDebug {
							fmt.Println("used long at +1", "better by", lengthLong-length-1)
						}
						// Override the match, since this is clearly better.
						length = lengthLong
						shortOff = longOff
						srci++
					}
				}
			}
			// Emit any pending literals
			emitLits()

			if printDebug {
				fmt.Println("Emitting short offset", srci-shortOff, "length", length)
			}
			// Matches have the lowest bit set to 1
			l := uint64(length)<<1 + 1
			dst = binary.AppendUvarint(dst, l)
			dst = binary.AppendUvarint(dst, uint64(srci-shortOff))

			// Index added bytes
			index(srci, srci+length)
			srci += length
			nextEmit = srci
			continue
		}
		// Update tables
		longTab[cur] = srci
		shortTab[uint32(cur)] = srci

		// Move to the next
		srci++
	}
	// Emit any pending literals
	srci = len(src)
	if srci > nextEmit {
		emitLits()
	}
	return string(dst)
}

func Decode(s string) string {
	const printDebug = false
	dst := make([]byte, 0, len(s)*2)
	src := []byte(s)

	// Read until we have no more bytes.
	for len(src) > 0 {
		length, n := binary.Uvarint(src)
		if n <= 0 {
			panic("invalid uvarint for length")
		}
		src = src[n:]

		// Lower bit is 0 for literals, 1 for matches
		if length&1 == 0 {
			// Discard the lowest bit
			litLen := length >> 1
			if printDebug {
				fmt.Println("Found", litLen, "literals")
			}
			// Append literals to output
			dst = append(dst, src[:litLen]...)
			src = src[litLen:]
			continue
		}

		// Discard the lowest bit
		ml := length >> 1
		matchLength := int(ml)

		// Read match offset.
		matchOffset, n := binary.Uvarint(src)
		if printDebug {
			fmt.Println("Found Match, offset:", matchOffset, "length:", matchLength)
		}
		if n <= 0 {
			panic("invalid uvarint for match offset")
		}
		src = src[n:]
		copyFrom := len(dst) - int(matchOffset)

		// If we overlap, we need to copy bytes one by one.
		if int(matchOffset) <= matchLength {
			for i := 0; i < matchLength; i++ {
				dst = append(dst, dst[copyFrom+i])
			}
		} else {
			dst = append(dst, dst[copyFrom:copyFrom+matchLength]...)
		}
	}
	return string(dst)
}

// decodeMinimal is a minimal version of the decode function
// for those that prefer a more concise version.
func decodeMinimal(s string) string {
	src, dst := []byte(s), make([]byte, 0)

	readUvarInt := func() uint64 {
		v, n := binary.Uvarint(src)
		if n <= 0 {
			panic("invalid uvarint")
		}
		src = src[n:]
		return v
	}

	for len(src) > 0 {
		v := readUvarInt()
		length := v >> 1
		if v&1 == 0 {
			dst = append(dst, src[:length]...)
			src = src[length:]
			continue
		}
		copyFrom := len(dst) - int(readUvarInt())
		for i := 0; i < int(length); i++ {
			dst = append(dst, dst[copyFrom+i])
		}
	}
	return string(dst)
}

// matchLen returns the number of bytes that match at the beginning of a and b.
// a must be the shortest slice.
func matchLen(a, b []byte) int {
	// Number of bytes already matched.
	n := 0

	// Quick loop, comparing 8 bytes at a time.
	for len(a) >= 8 {
		match := binary.LittleEndian.Uint64(a) ^ binary.LittleEndian.Uint64(b)
		// If a[:8] == b[:8], then match will be 0.
		if match != 0 {
			// All bits that are the same will be 0s.
			// Find the index of the first bit that is different
			// and divide by 8 to get the number of bytes.
			return n + bits.TrailingZeros64(match)/8
		}
		a = a[8:]
		b = b[8:]
		n += 8
	}

	// Handle the last bytes
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return n + i
		}
	}
	return n
}
