package decimaltobinary

import (
	"math"
	"strconv"
)

func Tobinary(n float64) string {
	res := ``
	for 1 == 1 {
		n = 2 * n
		res += strconv.Itoa(int(n))
		n = n - float64(int(n))
		if n == 0 {
			break
		}
	}
	return res
}

func Toten(b string) (r float64) {
	l := len(b)
	for i := 0; i < l; i++ {
		r += float64(b[i]-48) * math.Pow(2, -float64(i+1))
	}
	return
}
