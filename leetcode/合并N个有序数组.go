package main

import "fmt"

//88题的变种

func mergeN(numsList [][]int) []int {
	if len(numsList) == 0 {
		return []int{}
	}
	if len(numsList) == 1 {
		return numsList[0]
	}
	//方法一：归并排序，，但时间复杂度高，他适合无序数组排序
	//方法二：双指针, 正向 利用两个数组已经有序，双指针 正向 需要空间
	return mergeN_doublePointer_1(numsList)
	//方法三：双指针, 逆向 利用两个数组已经有序，双指针 反向 不需要空间，利用nums1后半部分为空的特性，把nums1当成方法二的临时数组用，参考88题
}

func mergeN_doublePointer_1(numsList [][]int) []int {
	preNums := numsList[0]
	for i := 1; i < len(numsList); i++ {
		curNums := numsList[i]
		tmp := make([]int, 0, len(preNums)+len(curNums))
		index1, index2 := 0, 0
		for {
			if index1 == len(preNums) {
				tmp = append(tmp, curNums[index2:]...)
				break
			}
			if index2 == len(curNums) {
				tmp = append(tmp, preNums[index1:]...)
				break
			}
			if preNums[index1] < curNums[index2] {
				tmp = append(tmp, preNums[index1])
				index1++
			} else {
				tmp = append(tmp, curNums[index2])
				index2++
			}
		}
		preNums = tmp
	}
	return preNums
}

func mergeN_doublePointer_2(nums1 []int, m int, nums2 []int, n int) {
	index1, index2, tailIndex := m-1, n-1, len(nums1)-1
	for index1 >= 0 || index2 >= 0 {
		if index1 == -1 {
			nums1[tailIndex] = nums2[index2]
			index2--
		} else if index2 == -1 {
			nums1[tailIndex] = nums1[index1]
			index1--
		} else if nums1[index1] > nums2[index2] {
			nums1[tailIndex] = nums1[index1]
			index1--
		} else {
			nums1[tailIndex] = nums2[index2]
			index2--
		}
		tailIndex--
	}
}

func main() {
	fmt.Println(mergeN([][]int{{1, 2, 4}, {2, 5, 6}, {1, 3, 5, 7}}))
}
