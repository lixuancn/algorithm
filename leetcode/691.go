package main

import (
	"fmt"
	"strings"
)

//691. 贴纸拼词

//回溯法
//拿target[0]看，这个sticker是否有这个字母，如果没有就跳过。因为target[0]总要被找到，所以并不会影响结果。
func minStickers(stickers []string, target string) int {
	result := 1 << 31
	cache := make(map[string]int, 0) //找到某个字符串，所用的最小贴纸数
	num := 0                         //已使用的贴纸数
	var backtracking func(stickers []string, left string)
	backtracking = func(stickers []string, left string) {
		//终止条件
		if left == "" {
			result = num
			return
		}
		//提前结束，因为下面num++后肯定就比result大了，不符合最小值的答案
		if num+1 >= result {
			return
		}
		for _, sticker := range stickers {
			if strings.IndexByte(sticker, left[0]) == -1 {
				continue
			}
			s1, s2 := sticker, left
			//删除left中能被ticker贴出来的字母
			for i := 0; i < len(left); i++ {
				pos := strings.IndexByte(s1, left[i])
				if pos != -1 {
					s1 = s1[:pos] + s1[pos+1:]
					left = left[:i] + left[i+1:]
					i--
				}
			}
			//贴纸数加一
			num++
			if v, ok := cache[left]; !ok || v > num {
				cache[left] = num
				backtracking(stickers, left)
			}
			//回溯到原状态
			num--
			left = s2
		}
	}
	backtracking(stickers, target)
	if result == 1<<31 {
		return -1
	}
	return result
}

func main() {
	fmt.Println(minStickers([]string{"with", "example", "science"}, "thehat")) //3
	fmt.Println(minStickers([]string{"notice", "possible"}, "basicbasic"))     //-1
}
