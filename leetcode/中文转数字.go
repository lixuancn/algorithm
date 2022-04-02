package main

import (
	"fmt"
)

func ChineseToNumber(s string) int64 {
	//ret是最终返回值，tmpRet是每次累加值，
	//如果遇到万或者亿会把值放到ret里并将tmpRet清空
	var ret, tmpRet, number, pos int64
	ret, number = 0, 0
	l := len(s)
	for k, v := range s {
		vNumber := GetNumber(string(v))
		if vNumber > 0 {
			number = vNumber
			if k == l-3 {
				tmpRet += number
				break
			}
		} else {
			pos = GetPosNumber(string(v))
			if string(v) == "万" || string(v) == "亿" {
				if string(v) == "亿" {
					ret = (ret + tmpRet + number) * pos
				} else {
					tmpRet = (tmpRet + number) * pos
					ret += tmpRet
				}
				tmpRet, number = 0, 0
			} else {
				tmpRet += number * pos
			}
			number = 0
		}
	}
	if tmpRet != 0 && pos != 10000 && pos != 100000000 {
		ret += tmpRet
	}
	return ret
}

func GetNumber(s string) int64 {
	var ret int64
	switch s {
	case "零":
		ret = 0
	case "一":
		ret = 1
	case "二":
		ret = 2
	case "三":
		ret = 3
	case "四":
		ret = 4
	case "五":
		ret = 5
	case "六":
		ret = 6
	case "七":
		ret = 7
	case "八":
		ret = 8
	case "九":
		ret = 9
	default:
		ret = -1
	}
	return ret
}

func GetPosNumber(s string) int64 {
	var ret int64
	switch s {
	case "个":
		ret = 1
	case "十":
		ret = 10
	case "百":
		ret = 100
	case "千":
		ret = 1000
	case "万":
		ret = 10000
	case "亿":
		ret = 100000000
	default:
		ret = -1
	}
	return ret
}

func main() {
	s := "四千九百四十二万五千三百六十八亿七千九百六十万零二百三十八"
	fmt.Println(ChineseToNumber(s))
}
