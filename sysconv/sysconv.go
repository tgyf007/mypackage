package sysconv

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var sixteenMap = map[string]int{`A`: 10, `B`: 11, `C`: 12, `D`: 13, `E`: 14, `F`: 15}

func Convert(s string, from, to int) (ret string, err error) {
	if len(s) == 0 || from == 0 || to == 0 {
		ret = ``
		return
	}

	var re int
	if from != 10 {
		re, err = convert(s, from)
		if err != nil {
			return
		}
	} else {
		re, _ = strconv.Atoi(s)
	}

	switch to {
	case 2:
		ret = fmt.Sprintf("%b", re)
		break
	case 8:
		ret = fmt.Sprintf("%o", re)
		break
	case 16:
		ret = fmt.Sprintf("%X", re)
		break
	default:
		ret = string(re)
	}
	return
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
