package main

import "fmt"

//面试题 17.11. 单词距离

//这里是一次问题。
//如果上层会多次调用本函数，即多次遍历这个文件的话，可以维护一个哈希表，key是每个单词，value是个链表，记录每个单词每次出现的位置。
//在查找时，哈希表选出2个链表，双指针开始遍历即可。
func findClosest(words []string, word1 string, word2 string) int {
	lastIndex1, lastIndex2 := -1, -1 //两个词最后一次出现的下标
	result := 1 << 31
	for i := 0; i < len(words); i++ {
		if word1 != words[i] && word2 != words[i] {
			continue
		}
		if word1 == words[i] {
			lastIndex1 = i
		} else if word2 == words[i] {
			lastIndex2 = i
		}
		if lastIndex1 >= 0 && lastIndex2 >= 0 {
			d := lastIndex2 - lastIndex1
			if d < 0 {
				d = -d
			}
			if result > d {
				result = d
			}
			if result == 1 {
				break
			}
		}
	}
	return result
}

func main() {
	fmt.Println(findClosest([]string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}, "a", "student"))
}
