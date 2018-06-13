package base64

import (
	"github.com/tgyf007/mypackage/binarytostring"
	"strconv"
	"strings"
)

const (
	BASE_LEN = 3
	NEW_LEN  = 6
	OLD_LEN  = 8
)

var valMap = map[int]string{
	0: `A`, 1: `B`, 2: `C`, 3: `D`, 4: `E`, 5: `F`, 6: `G`, 7: `H`, 8: `I`, 9: `J`,
	10: `K`, 11: `L`, 12: `M`, 13: `N`, 14: `O`, 15: `P`, 16: `Q`, 17: `R`, 18: `S`, 19: `T`,
	20: `U`, 21: `V`, 22: `W`, 23: `X`, 24: `Y`, 25: `Z`, 26: `a`, 27: `b`, 28: `c`, 29: `d`,
	30: `e`, 31: `f`, 32: `g`, 33: `h`, 34: `i`, 35: `j`, 36: `k`, 37: `l`, 38: `m`, 39: `n`,
	40: `o`, 41: `p`, 42: `q`, 43: `r`, 44: `s`, 45: `t`, 46: `u`, 47: `v`, 48: `w`, 49: `x`,
	50: `y`, 51: `z`, 52: `0`, 53: `1`, 54: `2`, 55: `3`, 56: `4`, 57: `5`, 58: `6`, 59: `7`,
	60: `8`, 61: `9`, 62: `+`, 63: `/`}

func Encode(str string) (ret string, err error) {
	str_len := len(str)
	if str_len == 0 {
		return
	}

	binstr, fill_str, fill_equal := ``, ``, ``
	for i := 0; i < str_len; i++ {
		binstr += fillNum(strconv.FormatInt(int64(str[i]), 2), OLD_LEN)
	}
	remain := 3 - str_len%3
	if remain != 3 {
		for i := 0; i < remain; i++ {
			fill_str += fillZero(OLD_LEN)
			fill_equal += `=`
		}
		binstr += fill_str
		binstr = binstr[0 : len(binstr)-remain*NEW_LEN]
	}
	splitBin := binarytostring.SplitByLen(binstr, NEW_LEN)
	res := buildNewResult(splitBin)
	ret = res + fill_equal
	return
}

func Decode(str string) (ret string, err error) {
	str_len := len(str)
	if str_len == 0 {
		return
	}
	equal_count := strings.Count(str, `=`)
	str = str[0 : str_len-equal_count]
	flipMap := make(map[string]int, len(valMap))
	for key, val := range valMap {
		flipMap[val] = key
	}
	binstr := ``
	new_len := str_len - equal_count
	for i := 0; i < new_len; i++ {
		binstr += fillNum(strconv.FormatInt(int64(flipMap[string(str[i])]), 2), NEW_LEN)
	}
	if equal_count > 0 {
		for i := 0; i < equal_count; i++ {
			binstr += fillZero(NEW_LEN)
		}
		binstr = binstr[0 : len(binstr)-equal_count*OLD_LEN]
	}
	ret = binarytostring.BtoS(binstr)
	return
}

func fillNum(ori string, num int) (res string) {
	ori_len := len(ori)
	if ori_len == num {
		res = ori
		return
	} else if ori_len > num {
		res = ``
		return
	} else {
		pre := ``
		for i := num; i > ori_len; i-- {
			pre += `0`
		}
		res = pre + ori
		return
	}
}

func fillZero(num int) (res string) {
	for i := 0; i < num; i++ {
		res += `0`
	}
	return
}

func buildNewResult(from []string) (res string) {
	from_len := len(from)
	mapV := valMap
	for i := 0; i < from_len; i++ {
		newCode, _ := strconv.ParseInt(from[i], 2, 0)
		res += mapV[int(newCode)]
	}
	return
}
