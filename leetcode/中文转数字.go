package main

import (
	"container/list"
	"fmt"
)

//未完 TODO

var chineseList = map[string]int{"零": 0, "一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9}
var unitList = map[string]int{"十": 10, "百": 100, "千": 1000, "万": 10000, "亿": 100000000}

func main() {
	fmt.Println(transformation("一千八百万"))
}

func transformation(str string) int {
	num := 0
	digint := 0
	//数字栈
	digintStack := list.New()
	//单位栈
	unitStack := list.New()
	for _, c := range str {
		s := string(c)
		//是数字，但其实中文语法只会有一个数字
		if _, ok := chineseList[s]; ok {
			digint = digint*10 + chineseList[s]
		} else if _, ok := unitList[s]; ok{
			//如果是单位
			digintStack.PushBack(digint)

				unitStack.PushBack(unitList[s])
			}
		}else{
			return 0
		}
		fmt.Println(string(c))
	}

	return num
}
