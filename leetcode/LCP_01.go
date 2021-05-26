package main
import (
	"fmt"
)

func main() {
	r := game([]int{2,2,3}, []int{1,2,3})
	fmt.Println(r)
}

func game(guess []int, answer []int) int {
	count := 0
	for i := range guess{
		if guess[i] == answer[i]{
			count++
		}
	}
	return count
}
