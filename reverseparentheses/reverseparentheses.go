package reverseparentheses

func Reverse(s string) string {
	return string(reverseBytes([]byte(s)))
}

func reverseBytes(in []byte) []byte {
	for i := 0; i < len(in); i++ {
		r := in[i]
		if r == '(' {
			in = append(in[:i], reverseBytes(in[i+1:])...)
		}
		if r == ')' {
			simpleReverse(in[:i])
			return append(in[:i], in[i+1:]...)
		}
	}
	return in

}

func simpleReverse(in []byte) {
	i := len(in)
	for j := 0; j < i/2; j++ {
		x := in[j]
		in[j] = in[i-j-1]
		in[i-j-1] = x
	}
}
