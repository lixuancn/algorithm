package main

import (
	"container/list"
	"fmt"
	"strings"
	"unicode"
)

//591. 标签验证器

//栈+字符串遍历
//如果遇到<  那么：
//1、i+1是/，下一个>的位置是j，那么[i+2, j-1]是TAG_NAME，看和栈顶是否一致，出栈
//2、i+1是!，说明遇到cdata，看[i+2,i+7]是不是[CDATA[，然后找到下一个]]>，位置为j，[i+9,j-1]就是cdata，不用管
//3、i+1是大写字母，找到下一个>位置为j，[i+2,j-1]就是TAG_NAME，语法校验后入栈
//4、除此以外，如果不存在下一个字符，或下一个字符不符合上述情况，则不合法
func isValid(code string) bool {
	stack := list.New()
	for code != "" {
		if code[0] == '<' {
			if len(code) == 1 {
				return false
			} else if code[1] == '/' {
				//找到第一个>
				j := strings.IndexByte(code, '>')
				if j == -1 {
					return false
				}
				if stack.Len() == 0 || code[2:j] != stack.Back().Value.(string) {
					return false
				}
				stack.Remove(stack.Back())
				code = code[j+1:]
				if stack.Len() == 0 && code != "" {
					return false
				}
			} else if code[1] == '!' {
				if stack.Len() == 0 || len(code) < 9 || code[2:9] != "[CDATA[" {
					return false
				}
				j := strings.Index(code, "]]>")
				if j == -1 {
					return false
				}
				code = code[j+1:]
			} else {
				j := strings.IndexByte(code, '>')
				if j == -1 {
					return false
				}
				tagName := code[1:j]
				if tagName == "" || len(tagName) > 9 {
					return false
				}
				for _, c := range tagName {
					if !unicode.IsUpper(c) {
						return false
					}
				}
				stack.PushBack(tagName)
				code = code[j+1:]
			}
		} else {
			//TAG_CONTENT必须包含在标签内
			if stack.Len() == 0 {
				return false
			}
			code = code[1:]
		}
	}
	return stack.Len() == 0
}

func main() {
	fmt.Println(isValid("<DIV>This is the first line <![CDATA[<div>]]></DIV>"))                               //true
	fmt.Println(isValid("<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>"))                                //true
	fmt.Println(isValid("<A>  <B> </A>   </B>"))                                                              //false
	fmt.Println(isValid("<DIV>  div tag is not closed  <DIV>"))                                               //false
	fmt.Println(isValid("<DIV>  unmatched <  </DIV>"))                                                        //false
	fmt.Println(isValid("<DIV> closed tags with invalid tag name  <b>123</b> </DIV>"))                        //false
	fmt.Println(isValid("<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>")) //false
	fmt.Println(isValid("<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>"))                //false
}
