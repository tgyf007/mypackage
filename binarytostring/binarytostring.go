package binarytostring

import (
	"github.com/tgyf007/mypackage/sysconv"
	"strconv"
	"strings"
)

func SplitByLen(str string, lenNum int) (res []string) {
	str_len := len(str)
	res_len := str_len / lenNum
	for i := 0; i < res_len; i++ {
		start := i * lenNum
		res = append(res, str[start:start+lenNum])
	}
	return
}

func BtoS(binary string) (ret string) {
	splitBin := SplitByLen(binary, 8)
	spl_len := len(splitBin)
	for i := 0; i < spl_len; {
		//判断第一个字符是0(48) 还是1(49) 如果是1 说明不是ASCII码
		if splitBin[i][0] == 49 {
			bnum := strings.Index(splitBin[i], `0`)
			tmp1 := splitBin[i][bnum+1:]
			for j := 1; j < bnum; j++ {
				tmp1 += splitBin[i+j][2:]
			}
			newCode, _ := sysconv.Convert(tmp1, 2, 10)
			ret += newCode
			i += bnum
		} else {
			newCode, _ := strconv.ParseInt(splitBin[i], 2, 0)
			ret += string(newCode)
			i++
		}
	}
	return
}

func BuildString(by, num int) (res string) {
	for i := 0; i < num; i++ {
		res += string(by)
	}
	return
}
