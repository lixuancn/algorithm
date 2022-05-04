package main

import (
	"fmt"
	"strings"
)

//937. 重新排列日志文件

func reorderLogFiles(logs []string) []string {
	letterArr := make([]string, 0)
	digitArr := make([]string, 0)
	for _, log := range logs {
		one := strings.IndexByte(log, ' ') + 1
		//数字
		if log[one] >= 48 && log[one] <= 57 {
			digitArr = append(digitArr, log)
		} else {
			//字母
			letterArr = append(letterArr, log)
		}
	}
	//字母排序
	merge(letterArr, 0, len(letterArr)-1)
	return append(letterArr, digitArr...)
}

//归并
func merge(letterArr []string, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	merge(letterArr, left, mid)
	merge(letterArr, mid+1, right)
	mergeSort(letterArr, left, mid, right)
}

func mergeSort(letterArr []string, left, mid, right int) {
	tmp := make([]string, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if compare(letterArr[i], letterArr[j]) {
			tmp[k] = letterArr[i]
			i++
		} else {
			tmp[k] = letterArr[j]
			j++
		}
		k++
	}
	for ; i <= mid; i, k = i+1, k+1 {
		tmp[k] = letterArr[i]
	}
	for ; j <= right; j, k = j+1, k+1 {
		tmp[k] = letterArr[j]
	}
	for i := range tmp {
		letterArr[left+i] = tmp[i]
	}
}

//a小就true，b小就false
func compare(a, b string) bool {
	space1 := strings.IndexByte(a, ' ')
	space2 := strings.IndexByte(b, ' ')
	if a[space1:] < b[space2:] {
		return true
	} else if a[space1:] > b[space2:] {
		return false
	} else {
		return a[:space1] < b[:space2]
	}
}

func main() {
	fmt.Println(reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"})) //let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6
	fmt.Println(reorderLogFiles([]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}))          //"g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"

}
