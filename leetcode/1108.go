package main
import (
	"fmt"
)

//给你一个有效的 IPv4 地址 address，返回这个 IP 地址的无效化版本。
//所谓无效化 IP 地址，其实就是用  "[.]" 代替了每个 "."。

func main() {
	r := defangIPaddr("1.1.1.1")
	fmt.Println(r)
}

func defangIPaddr(address string) string {
	n := len(address)
	var ret = make([]byte, n+6)
	offset := 0
	for i := range address{
		key := offset + i
		if address[i] == '.'{
			ret[key] = '['
			ret[key+1] = '.'
			ret[key+2] = ']'
			offset += 2
		}else{
			ret[key] = address[i]
		}
	}
	return string(ret)
}
