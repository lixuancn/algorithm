package main

import "fmt"

func main() {
	r := intToRoman(1994)
	fmt.Println(r)
}

func intToRoman(num int) string {
	ruleList := [7]int{1000, 500, 100, 50, 10, 5, 1,}
	ruleMap := map[int]string{
		1000: "M",
		500: "D",
		100: "C",
		50: "L",
		10: "X",
		5: "V",
		1: "I",
	}
	ret := ""
	for _, v := range ruleList{
		count := num / v
		if num == 4{
			ret += "IV"
			num -= 4
			continue
		}else if num == 9{
			ret += "IX"
			num -= 9
			continue
		}else if num >= 40 && num <= 49{
			ret += "XL"
			num -= 40
			continue
		}else if num >= 90 && num <= 99{
			ret += "XC"
			num -= 90
			continue
		}else if num >= 400 && num <= 499{
			ret += "CD"
			num -= 400
			continue
		}else if num >= 900 && num <= 999{
			ret += "CM"
			num -= 900
			continue
		}

		if count <= 0{
			continue
		}
		for i:=0; i<count; i++{
			ret += ruleMap[v]
		}
		num = num - count * v
	}
	return ret
}
