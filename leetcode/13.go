package main

//13. 罗马数字转整数

func romanToInt(s string) int {
	result := 0
	symbolValues := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}
