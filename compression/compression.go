package compression

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"unicode/utf8"

	"encoding/binary"
)

// represents a node in a Hufmann Tree, holds character and it's number of occurances.
type HNode struct {
	char        rune
	freq        int
	left, right *HNode
}

// implements the interface Orderable
func (hn *HNode) Less(other Orderable) bool {
	otherhn, ok := other.(*HNode)
	if !ok {
		log.Printf("Type assertion failed: expected *HNode, got %T", other)
		return false
	}
	return hn.freq < otherhn.freq
}

// builds HNodes and puts them in a minimum priority queue
func buildHNodesPQ(chars []rune) *MinPQ[Orderable] {
	charFreq, hNodesPQ := make(map[rune]int), NewMinPQ[Orderable]()

	for _, r := range chars {
		charFreq[r] += 1
	}

	for k, v := range charFreq {
		hf := &HNode{
			char: k,
			freq: v,
		}
		hNodesPQ.Insert(hf)
	}

	return hNodesPQ
}

// builds a Hufmann Tree by picking the HNode having the least character frequency
func buildHuffmannTree(hfNodesPQ *MinPQ[Orderable]) *HNode {
	var a, b *HNode
	for !hfNodesPQ.IsEmpty() {
		temp, err := hfNodesPQ.DeleteMin()
		a = temp.(*HNode)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if hfNodesPQ.IsEmpty() {
			break
		} else {
			temp, err := hfNodesPQ.DeleteMin()
			b = temp.(*HNode)
			if err != nil {
				log.Fatalln(err.Error())
			}
		}
		c := &HNode{
			freq:  a.freq + b.freq,
			left:  a,
			right: b,
		}
		hfNodesPQ.Insert(c)
	}

	return a
}

// builds the mapping to encode the data. CodeCharMap maps the character to the Hufmann Code generated using the Hufmann Tree
func buildCodeCharMap(root *HNode) map[string]rune {
	codeCharMap := make(map[string]rune)
	var buildCodes func(node *HNode, codeBuffer *bytes.Buffer)
	buildCodes = func(node *HNode, codeBuffer *bytes.Buffer) {
		if node.left == nil && node.right == nil {
			codeCharMap[codeBuffer.String()] = node.char
			return
		}
		if node.left != nil {
			codeBuffer.WriteByte('0')
			buildCodes(node.left, codeBuffer)
			codeBuffer.Truncate(codeBuffer.Len() - 1)
		}
		if node.right != nil {
			codeBuffer.WriteByte('1')
			buildCodes(node.right, codeBuffer)
			codeBuffer.Truncate(codeBuffer.Len() - 1)
		}
	}

	if root.left == nil && root.right == nil {
		codeCharMap["0"] = root.char
		return codeCharMap
	}
	var codeBuffer bytes.Buffer
	buildCodes(root, &codeBuffer)
	return codeCharMap
}

// Serializes the mapping used for encoding the data. Formatting: for each character and it's Hufmann code, <code length><characterlength><code><character>
func serializeCodeCharMap(codeCharMap map[string]rune) ([]byte, error) {
	var encodedBytesBuffer bytes.Buffer

	for k, v := range codeCharMap {
		keyLen := len(k)
		valLen := utf8.RuneLen(v)
		if valLen == -1 {
			return nil, fmt.Errorf("invalid UTF-8 rune: %v", v)
		}

		// write key length and value length as single bytes
		encodedBytesBuffer.WriteByte(byte(keyLen))
		encodedBytesBuffer.WriteByte(byte(valLen))

		// write the key (code) and character
		encodedBytesBuffer.WriteString(k)
		encodedBytesBuffer.WriteRune(v)
	}

	return encodedBytesBuffer.Bytes(), nil
}

// Reconstructs the mapping used for encoding the data.
func deserializeCodeCharMap(encodedBytesBuffer *bytes.Buffer, encodingMappingLen int32) (map[string]rune, error) {
	codeCharMap := make(map[string]rune)
	n := 0

	for n < int(encodingMappingLen) {
		// Read the key length (uint8)
		keyLenByte, err := encodedBytesBuffer.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("failed to read key length: %v", err)
		}
		n += 1
		keyLen := int(keyLenByte)

		// Read the value length (uint8)
		valLenByte, err := encodedBytesBuffer.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("failed to read value length: %v", err)
		}
		n += 1
		valLen := int(valLenByte)

		// Read the key
		key := make([]byte, keyLen)
		if _, err := encodedBytesBuffer.Read(key); err != nil {
			return nil, fmt.Errorf("failed to read key: %v", err)
		}
		n += keyLen
		// Read the value (should be one rune)
		valueBytes := make([]byte, valLen)
		if _, err := encodedBytesBuffer.Read(valueBytes); err != nil {
			return nil, fmt.Errorf("failed to read value: %v", err)
		}
		value, size := utf8.DecodeRune(valueBytes)
		if size == 0 || value == utf8.RuneError {
			log.Printf("Warning: invalid UTF-8 rune detected at position %d", n)
			return nil, fmt.Errorf("invalid UTF-8 rune detected")
		}
		n += valLen
		// Add to map
		codeCharMap[string(key)] = value
	}

	return codeCharMap, nil
}

// encode the given string returns a string in the format: <length of encoding mapping><encoding mapping><encoded string>
func Encode(s string) string {
	chars := []rune(s)
	pq := buildHNodesPQ(chars)
	root := buildHuffmannTree(pq)

	// codeCharMap will be used for decoding the string
	codeCharMap := buildCodeCharMap(root)

	var encodedBytesBuffer bytes.Buffer
	encodingMappingBytes, err := serializeCodeCharMap(codeCharMap)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// writing the length of the mapping used for encoding the given string s
	err = binary.Write(&encodedBytesBuffer, binary.BigEndian, int32(len(encodingMappingBytes)))
	if err != nil {
		log.Fatalln("Error occured in writing the length of the encoding mapping to the buffer:", err)
	}
	// writing the encoding mapping itself
	encodedBytesBuffer.Write(encodingMappingBytes)

	// charCodeMap will be used for encoding the string
	charCodeMap := make(map[rune]string)
	for k, v := range codeCharMap {
		charCodeMap[v] = k
	}

	for _, el := range chars {
		encodedBytesBuffer.WriteString(charCodeMap[el])
	}

	return encodedBytesBuffer.String()
}

// decode the given string
func Decode(s string) string {
	encodedDataBytes := bytes.NewBuffer([]byte(s))
	var encodingMappingLen int32
	err := binary.Read(encodedDataBytes, binary.BigEndian, &encodingMappingLen)
	if err != nil {
		log.Fatalln("Error reading encoding mapping length:", err)
	}

	codeCharMap, err := deserializeCodeCharMap(encodedDataBytes, int32(encodingMappingLen))
	if err != nil {
		log.Fatalln(err.Error())
	}

	var sofar bytes.Buffer
	var decodedStringBuilder strings.Builder

	for encodedDataBytes.Len() > 0 {
		k, err := encodedDataBytes.ReadByte()
		if err != nil {
			log.Fatalln(err.Error())
		}
		sofar.WriteByte(k)
		if v, ok := codeCharMap[sofar.String()]; ok {
			decodedStringBuilder.WriteRune(v)
			sofar.Reset()
		}
	}
	return decodedStringBuilder.String()
}
