package chess

import "errors"

// CanKnightAttack checks if two knights on specified positions can attack each other
func CanKnightAttack(white, black string) (bool, error) {
	// Do initial validation
	if len(white) < 2 || len(black) < 2 {
		return false, errors.New("args too short")
	}

	// Check if pieces are on board
	if white[0]-'a' > 7 || white[1]-'1' > 7 {
		return false, errors.New("Invalid white position")
	}
	if black[0]-'a' > 7 || black[1]-'1' > 7 {
		return false, errors.New("Invalid black position")
	}

	// Calculate distance between knights in each axis
	d0 := abs(int(white[0]) - int(black[0]))
	d1 := abs(int(white[1]) - int(black[1]))

	// If the difference is 1 on one axis and 2 on the other we're happy
	if (d0 == 1 && d1 == 2) || (d0 == 2 && d1 == 1) {
		return true, nil
	} else if d0 == 0 && d1 == 0 {
		return false, errors.New("You can't stack 'em you silly.")
	}
	return false, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
