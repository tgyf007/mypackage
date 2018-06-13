package sysconv

import (
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var sixteenMap = map[string]int{`A`: 10, `B`: 11, `C`: 12, `D`: 13, `E`: 14, `F`: 15}

func Convert() {
	if len(os.Args) < 4 {
		return
	}
	s := os.Args[1]
	fromType, _ := strconv.Atoi(os.Args[2])
	toType := os.Args[3]

	var re int
	var err error
	if fromType != 10 {
		re, err = convert(s, fromType)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		re, _ = strconv.Atoi(s)
	}

	switch toType {
	case `2`:
		fmt.Printf("%b", re)
		break
	case `8`:
		fmt.Printf("%o", re)
		break
	case `16`:
		fmt.Printf("%X", re)
		break
	default:
		fmt.Println(re)
	}

}

func convert(v string, conType int) (re int, err error) {
	length := len(v)
	if length == 0 {
		re = 0
		return
	}
	b, _ := regexp.MatchString(`^[0-9A-Fa-f]+$`, v)
	if !b {
		err = errors.New(`wring value`)
		return
	}
	var res float64
	res = 0
	for i := 1; i <= length; i++ {
		tmpV1 := strings.ToUpper(string(v[i-1]))
		tmpV, ok := sixteenMap[tmpV1]
		if !ok {
			tmpV, _ = strconv.Atoi(string(tmpV1))
		}
		res += math.Pow(float64(conType), float64(length-i)) * float64(tmpV)
	}
	sre := strconv.FormatFloat(res, 'f', 0, 64)
	re, _ = strconv.Atoi(sre)
	return
}
