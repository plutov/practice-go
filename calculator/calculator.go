package calculator

import (
	"errors"
	"strconv"
	"strings"
)

// An AlgebraicSymbol is one of the four algebraic operators, or a bracket.
type AlgebraicSymbol int

const (
	Plus = iota
	Minus
	Multiply
	Divide
	LBracket
	RBracket
)

// A Token is either a literal or an algebraic symbol.
type Token struct {
	isSymbol bool
	literal  float64
	symbol   AlgebraicSymbol
}

// Convert an expression from infix notation to Reverse Polish notation (RPN).
// This is an implementation of the shunting yard algorithm ( https://en.wikipedia.org/wiki/Shunting_yard_algorithm ).
func infixToRPN(tokens []Token) []Token {
	output := make([]Token, 0)
	operatorStack := make([]Token, 0)

	for _, token := range tokens {
		if token.isSymbol {
			for len(operatorStack) > 0 && token.symbol <= operatorStack[len(operatorStack)-1].symbol {
				output = append(output, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = append(operatorStack, token)
		} else {
			output = append(output, token)
		}
	}

	for i := len(operatorStack) - 1; i >= 0; i-- {
		output = append(output, operatorStack[i])
	}

	return output
}

// Evaluate an expression which is guaranteed to contain no brackets.
func EvalNoBrackets(tokens []Token) (float64, error) {
	tokensRPN := infixToRPN(tokens)

	// Evaluate an expression in RPN.
	result := make([]float64, 0)
	for _, token := range tokensRPN {
		if token.isSymbol {
			// The unary minus operator is a special case.
			if len(result) == 1 && token.symbol == Minus {
				result[0] *= -1
				continue
			}

			if len(result) < 2 {
				return 0, errors.New("Not enough literals on the stack to perform a binary computation.")
			}

			val2 := result[len(result)-1]
			val1 := result[len(result)-2]
			var valResult float64

			switch token.symbol {
			case Plus:
				valResult = val1 + val2
			case Minus:
				valResult = val1 - val2
			case Multiply:
				valResult = val1 * val2
			case Divide:
				valResult = val1 / val2
			default:
				return 0, errors.New("Unexpected symbol.")
			}

			result = append(result[:len(result)-2], valResult)
		} else {
			result = append(result, token.literal)
		}
	}

	if len(result) != 1 {
		return 0, errors.New("Unexpected number of literals left on the stack after performing all computations.")
	}

	return result[0], nil
}

// Tokenise the input string.
func Tokenise(expr string) ([]Token, error) {
	stringToAlgebraicSymbol := map[byte]AlgebraicSymbol{
		'+': Plus,
		'-': Minus,
		'*': Multiply,
		'/': Divide,
		'(': LBracket,
		')': RBracket,
	}

	expr = strings.ReplaceAll(expr, " ", "")
	tokens := make([]Token, 0)
	firstUnparsedIndex := 0
	for i := range expr {
		symbol, isSymbol := stringToAlgebraicSymbol[expr[i]]
		if isSymbol {
			if firstUnparsedIndex != i {
				literalVal, err := strconv.ParseFloat(expr[firstUnparsedIndex:i], 64)
				if err != nil {
					return []Token{}, err
				}
				tokens = append(tokens, Token{false, literalVal, 0})
			}
			firstUnparsedIndex = i + 1
			tokens = append(tokens, Token{true, 0, symbol})
		}
	}
	if firstUnparsedIndex < len(expr) {
		literalVal, err := strconv.ParseFloat(expr[firstUnparsedIndex:], 64)
		if err != nil {
			return []Token{}, err
		}
		tokens = append(tokens, Token{false, literalVal, 0})
	}

	return tokens, nil
}

func Eval(expr string) (float64, error) {
	tokens, err := Tokenise(expr)
	if err != nil {
		return 0, err
	}

	// The level of nesting is increased by one every time a bracket is opened.
	nestedExprs := make([][]Token, 0)
	nestedExprs = append(nestedExprs, make([]Token, 0))
	nestingLevel := 0

	for _, token := range tokens {
		if token.isSymbol && token.symbol == LBracket {
			// Increase the nesting level by one.
			nestedExprs = append(nestedExprs, make([]Token, 0))
			nestingLevel++
		} else if token.isSymbol && token.symbol == RBracket {
			if nestingLevel == 0 {
				return 0, errors.New("Too many closing brackets.")
			}
			// The highest nesting level has been closed, so we can evaluate its expression and record the result in the next level of nesting.
			result, err := EvalNoBrackets(nestedExprs[nestingLevel])
			if err != nil {
				return 0, err
			}
			nestedExprs = nestedExprs[:nestingLevel]
			nestingLevel--
			nestedExprs[nestingLevel] = append(nestedExprs[nestingLevel], Token{false, result, 0})
		} else {
			nestedExprs[nestingLevel] = append(nestedExprs[nestingLevel], token)
		}
	}

	if nestingLevel != 0 || len(nestedExprs) != 1 {
		return 0, errors.New("Too many opening brackets.")
	}

	result, err := EvalNoBrackets(nestedExprs[nestingLevel])
	if err != nil {
		return 0, err
	}

	return result, nil
}
