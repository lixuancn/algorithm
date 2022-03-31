package main

import (
	"container/list"
	"fmt"
)

func checkValidString(s string) bool {
	//return checkValidString_stack(s)
	return checkValidString_doubleIterator(s)
	return checkValidString_tanxin(s)
}

//栈
func checkValidString_stack(s string) bool {
	leftStack := list.New()
	symbolStack := list.New()
	//用右括号来抵消左括号和*
	for i, c := range s {
		if c == '(' {
			leftStack.PushBack(i)
		} else if c == '*' {
			symbolStack.PushBack(i)
		} else if leftStack.Len() > 0 {
			leftStack.Remove(leftStack.Back())
		} else if symbolStack.Len() > 0 {
			symbolStack.Remove(symbolStack.Back())
		} else {
			return false
		}
	}
	//只剩左括号和*的话互相抵消，但是*需要充当)，因此*的位置必须在(的右边
	for leftStack.Len() > 0 && symbolStack.Len() > 0 {
		if leftStack.Remove(leftStack.Back()).(int) > symbolStack.Remove(symbolStack.Back()).(int) {
			return false
		}
	}

	//因为*可以视为空，所以不需要判断symbolStack.Len()
	return leftStack.Len() == 0
}

//两次遍历，从左往右确定)，保证()), *))*这样的。 再从右往左确定(，保证没有)((，*))*
func checkValidString_doubleIterator(s string) bool {
	l, r, m := 0, 0, 0
	for _, c := range s {
		if c == '(' {
			l++
		} else if c == ')' {
			l--
		} else {
			m++
		}
		//左不够，用*顶替
		if l < 0 {
			m--
			l++
		}
		if m < 0 {
			return false
		}
	}
	m = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			r++
		} else if s[i] == '(' {
			r--
		} else {
			m++
		}
		if r < 0 {
			r++
			m--
		}
		if m < 0 {
			return false
		}
	}
	return true
}

//贪心，遇到左就加一，右就减一，*可能加可能减，所以维护一个最大值和最小值，最小值减，最大值加。但最小值不能小于0
func checkValidString_tanxin(s string) bool {
	minLeft, maxLeft := 0, 0
	for _, c := range s {
		if c == '(' {
			minLeft++
			maxLeft++
		} else if c == ')' {
			minLeft--
			maxLeft--
		} else {
			minLeft--
			maxLeft++
		}
		if minLeft < 0 {
			minLeft = 0
		}
		if maxLeft < 0 {
			return false
		}
	}
	return minLeft == 0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(checkValidString("(*))"))
}
