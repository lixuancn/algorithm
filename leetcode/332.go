package main

import (
	"fmt"
	"sort"
)

var d map[string][]string
var ans []string

//挺难理解
//https://leetcode-cn.com/problems/reconstruct-itinerary/solution/332-zhong-xin-an-pai-xing-cheng-shen-sou-hui-su-by/

func findItinerary(tickets [][]string) []string {
	//整理我手里的所有票，从v[0]出发，可以到的所有地方是哪里
	d = map[string][]string{}
	for _, v := range tickets {
		d[v[0]] = append(d[v[0]], v[1])
	}
	for _, v := range d {
		sort.Strings(v)
	}
	fmt.Println(d)
	ans = []string{}
	dfs("JFK")
	return ans
}

func dfs(from string) {
	for len(d[from]) > 0 {
		v := d[from][0]
		d[from] = d[from][1:]
		dfs(v)
	}
	ans = append([]string{from}, ans...)
}

func main() {
	//fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
}
