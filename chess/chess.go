package chess

import (
	"errors"
	"math"
)

var invalidSquareErr = errors.New("invalid square")

// Knights can attack when seperated by 2 squares in one direction and 1 square in the other direction.
// Horizontal axis is called 'file' labelled from 'a' to 'f'
// Forward axis is called 'rank' labelled from 1 to 8
func CanKnightAttack(white, black string) (bool, error) {
	// square must have 2 characters
	if len(white) != 2 || len(black) != 2 {
		return true, invalidSquareErr
	}

	// cannot be on the same square
	if white == black {
		return false, invalidSquareErr
	}

	// assign integer values (1 - 8) to file letters and rank digits
	wFile := int(white[0] - 'a' + 1)
	if wFile < 1 || wFile > 8 {
		return false, invalidSquareErr
	}

	bFile := int(black[0] - 'a' + 1)
	if bFile < 1 || bFile > 8 {
		return false, invalidSquareErr
	}

	wRank := int(white[1] - '1' + 1)
	if wRank < 1 || wRank > 8 {
		return false, invalidSquareErr
	}

	bRank := int(black[1] - '1' + 1)
	if bRank < 1 || bRank > 8 {
		return false, invalidSquareErr
	}

	// check file and rank relative positions
	switch int(math.Abs(float64(wFile - bFile))) {
	case 1:
		return int(math.Abs(float64(wRank-bRank))) == 2, nil
	case 2:
		return int(math.Abs(float64(wRank-bRank))) == 1, nil
	}

	// white and black valid but not in knight range of each other
	return false, nil
}
