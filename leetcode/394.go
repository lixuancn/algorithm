package main

import (
	"container/list"
	"strconv"
	"strings"
)

//394. 字符串解码
func decodeString(s string) string {
	ret := ""
	charStack := list.New()
	multipleStack := list.New()
	multiple := 0
	for _, c := range s {
		//先取倍数
		digit, err := strconv.Atoi(string(c))
		if err == nil {
			multiple = multiple*10 + digit
		} else if c == '[' {
			multipleStack.PushBack(multiple)
			multiple = 0
			charStack.PushBack(ret)
			ret = ""
		} else if c == ']' {
			//题目保证了一定有效，所以不用判断stack.Len==0的情况
			count := multipleStack.Remove(multipleStack.Back()).(int)
			str := charStack.Remove(charStack.Back()).(string)
			ret = str + strings.Repeat(ret, count)
		} else {
			ret += string(c)
		}
	}
	return ret
}

func main() {
	//fmt.Println(decodeString("3[a]2[bc]"))
	//fmt.Println(decodeString("3[a2[c]]"))
	//fmt.Println(decodeString("2[abc]3[cd]ef"))
}
