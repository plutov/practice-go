package functionfrequency

import (
	"sort"
)

const (
	eventLetter = iota
	eventUnderscore
	eventNumber
	eventDoubleQuote
	eventApostrophe
	eventBackSlesh
	eventNewLine
	eventTabSpace
	eventStartParenthesis
	eventPoint
	eventOther
)

const (
	stateStrtLine = iota // start line
	stateSkipLine        // skip line state
	stateContinue        // continue
	stateSkipWord        // skip word
	stateIgNxSymb        // ignore next symbol
	stateIsString        // is string
	stateStrIgnor        // string ignore
	stateStartFun        // start function
	stateMidleFun        // middle function
	stateMidPtFun        // middle point function
	stateEndinFun        // end function
)

var states = [][]int{
	// state/event  letter         underscore     number         double quote   apostrophe     back slesh     new line       tab/space      start parenth  point          other
	stateStrtLine: {stateSkipLine, stateSkipLine, stateSkipLine, stateSkipLine, stateSkipLine, stateSkipLine, stateStrtLine, stateContinue, stateSkipLine, stateSkipLine, stateSkipLine},
	stateSkipLine: {stateSkipLine, stateSkipLine, stateSkipLine, stateSkipLine, stateSkipLine, stateSkipLine, stateStrtLine, stateSkipLine, stateSkipLine, stateSkipLine, stateSkipLine},
	stateContinue: {stateStartFun, stateStartFun, stateSkipWord, stateIsString, stateIgNxSymb, stateIgNxSymb, stateStrtLine, stateContinue, stateContinue, stateSkipWord, stateSkipWord},
	stateSkipWord: {stateSkipWord, stateSkipWord, stateSkipWord, stateIsString, stateIgNxSymb, stateIgNxSymb, stateStrtLine, stateContinue, stateContinue, stateSkipWord, stateSkipWord},
	stateIgNxSymb: {stateContinue, stateContinue, stateContinue, stateContinue, stateContinue, stateIgNxSymb, stateContinue, stateContinue, stateContinue, stateContinue, stateContinue},
	stateIsString: {stateIsString, stateIsString, stateIsString, stateContinue, stateIsString, stateStrIgnor, stateIsString, stateIsString, stateIsString, stateIsString, stateIsString},
	stateStrIgnor: {stateIsString, stateIsString, stateIsString, stateIsString, stateIsString, stateIsString, stateIsString, stateIsString, stateIsString, stateIsString, stateIsString},
	stateStartFun: {stateMidleFun, stateMidleFun, stateMidleFun, stateIsString, stateIgNxSymb, stateIgNxSymb, stateStrtLine, stateContinue, stateEndinFun, stateMidleFun, stateMidleFun},
	stateMidleFun: {stateMidleFun, stateMidleFun, stateMidleFun, stateIsString, stateIgNxSymb, stateIgNxSymb, stateStrtLine, stateContinue, stateEndinFun, stateMidPtFun, stateMidleFun},
	stateMidPtFun: {stateMidleFun, stateMidleFun, stateMidleFun, stateIsString, stateIgNxSymb, stateIgNxSymb, stateStrtLine, stateContinue, stateSkipWord, stateContinue, stateMidleFun},
	stateEndinFun: {stateStartFun, stateStartFun, stateSkipWord, stateIsString, stateIgNxSymb, stateIgNxSymb, stateStrtLine, stateContinue, stateContinue, stateSkipWord, stateContinue},
}

type state struct {
	s int
}

func initState() state {
	return state{s: stateStrtLine}
}

func (s state) getState() int {
	return s.s
}

func (s state) newState(event int) state {
	return state{s: states[s.s][event]}
}

func getEvent(p byte) int {
	switch {
	case p >= 'a' && p <= 'z' || p >= 'A' && p <= 'Z':
		return eventLetter
	case p == '_':
		return eventLetter
	case p >= '0' && p <= '9':
		return eventNumber
	case p == '"':
		return eventDoubleQuote
	case p == '\'':
		return eventApostrophe
	case p == '\\':
		return eventBackSlesh
	case p == '\n':
		return eventNewLine
	case p == '\t' || p == ' ':
		return eventTabSpace
	case p == '(':
		return eventStartParenthesis
	case p == '.':
		return eventPoint
	}
	return eventOther
}

// FunctionFrequencyFSM returns the top 3 most mentioned functions in the code sample
func FunctionFrequencyFSM(gocode []byte) []string {
	functions := readFunctions(gocode)

	return topStrings(functions, 3)
}

func readFunctions(gocode []byte) map[string]int {
	functions := map[string]int{}

	s := initState()
	startIndex := -1
	for i := range gocode {
		s = s.newState(getEvent(gocode[i]))

		switch s.getState() {
		case stateStartFun:
			startIndex = i
		case stateEndinFun:
			if startIndex != -1 {
				function := string(gocode[startIndex:i])
				if function != "func" {
					functions[function]++
				}
				startIndex = -1
			}
		case stateMidleFun:
			continue
		case stateMidPtFun:
			continue
		default:
			startIndex = -1
		}

	}

	return functions
}

func topStrings(countedStrings map[string]int, top int) []string {
	strings := keysToSlice(countedStrings)

	sort.Slice(strings, func(i, j int) bool {
		return countedStrings[strings[i]] > countedStrings[strings[j]]
	})

	return strings[:top]
}

func keysToSlice(keys map[string]int) []string {
	slice := make([]string, 0, len(keys))
	for k := range keys {
		slice = append(slice, k)
	}
	return slice
}
