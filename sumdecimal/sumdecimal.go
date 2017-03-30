package sumdecimal

import (
	"math"
	"math/big"
)

func SumDecimal(c int) int {
	if c < 1 {
		return 0
	}

	sum := int64(0)

	p := int(math.Sqrt(float64(c)))

	c -= p * p

	bigC := big.NewInt(int64(c))
	bigX := big.NewInt(0)
	bigS := big.NewInt(0)
	bigM := big.NewInt(0)
	bigP := big.NewInt(int64(p * 20))

	one := big.NewInt(1)
	ten := big.NewInt(10)
	twenty := big.NewInt(20)
	hundred := big.NewInt(100)

	for i := 0; i < 1000; i++ {
		bigC.Mul(bigC, hundred)
		bigX.DivMod(bigC, bigP, bigM)
		bigS.Mul(bigX, bigX)

		bigC.Set(bigM)
		bigC.Sub(bigC, bigS)

		if bigS.Cmp(bigM) > 0 {
			bigX.Sub(bigX, one)
			bigC.Add(bigC, bigP)
			bigC.Add(bigC, bigX)
			bigC.Add(bigC, bigX)
			bigC.Add(bigC, one)
		}

		sum += bigX.Int64()

		bigX.Mul(bigX, twenty)
		bigP.Mul(bigP, ten)
		bigP.Add(bigP, bigX)
	}

	return int(sum)
}
