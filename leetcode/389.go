package main

func findTheDifference(s string, t string) byte {
	record := map[byte]int{}
	for i := range s {
		record[s[i]]++
	}
	for i := range t {
		record[t[i]]--
		if record[t[i]] < 0 {
			return t[i]
		}
	}
	return byte(' ')
}
