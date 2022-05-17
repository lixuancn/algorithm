package main

import "fmt"

//953. 验证外星语词典

func isAlienSorted(words []string, order string) bool {
	if len(words) < 2 {
		return true
	}
	charToDigint := [26]int{}
	for i := range order {
		charToDigint[order[i]-'a'] = i
	}
	for i := 1; i < len(words); i++ {
		word1 := words[i-1]
		word2 := words[i]
		for k := 0; k < len(word1) && k < len(word2); k++ {
			if charToDigint[word1[k]-'a'] < charToDigint[word2[k]-'a'] {
				break
			} else if charToDigint[word1[k]-'a'] > charToDigint[word2[k]-'a'] {
				return false
			}
			if (k == len(word1)-1 || k == len(word2)-1) && len(word1) > len(word2) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))    //true
	fmt.Println(isAlienSorted([]string{"kuvp", "q"}, "ngxlkthsjuoqcpavbfdermiywz"))            //true
	fmt.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz")) //false
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))         //false
}
