package main

import "fmt"

//4. 寻找两个正序数组的中位数

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//return findMedianSortedArrays_sort_merge(nums1, nums2)         //暴力法，先排序-归并
	//return findMedianSortedArrays_sort_doublePointer(nums1, nums2) //暴力法，先排序-不用额外空间
	return findMedianSortedArrays_mid(nums1, nums2) //二分
}

//方法一，先排序
//时间O(m+n)，空间O(m+n)
func findMedianSortedArrays_sort_merge(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	nums := make([]int, m+n)
	i, i1, i2 := 0, 0, 0
	for i1 < m && i2 < n {
		if nums1[i1] < nums2[i2] {
			nums[i] = nums1[i1]
			i1++
			i++
		} else {
			nums[i] = nums2[i2]
			i2++
			i++
		}
	}
	for i1 < m {
		nums[i] = nums1[i1]
		i1++
		i++
	}
	for i2 < n {
		nums[i] = nums2[i2]
		i2++
		i++
	}
	result := float64(nums[(m+n)/2])
	if (m+n)&1 == 0 {
		result = float64(nums[(m+n)/2-1]+nums[(m+n)/2]) / 2
	}
	return result
}

//方法二，用双指针来降低空间
//时间O(m+n)，空间O(1)
func findMedianSortedArrays_sort_doublePointer(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	i1, i2 := 0, 0
	leftNum, rightNum := 0, 0
	target1, target2 := 0, 0
	if (m+n)&1 == 1 {
		target1 = (m+n)/2 + 1
		target2 = (m+n)/2 + 1
	} else {
		target1 = (m + n) / 2
		target2 = (m+n)/2 + 1
	}

	for i1 < m || i2 < n {
		fmt.Println(i1, i2)
		num := 0
		if i1 == m && i2 != n {
			num = nums2[i2]
			i2++
		} else if i1 != m && i2 == n {
			num = nums1[i1]
			i1++
		} else if nums1[i1] < nums2[i2] {
			num = nums1[i1]
			i1++
		} else {
			num = nums2[i2]
			i2++
		}
		if i1+i2 == target1 {
			leftNum = num
		}
		if i1+i2 == target2 {
			rightNum = num
			break
		}
	}
	return float64(leftNum+rightNum) / 2
}

func findMedianSortedArrays_mid(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if (m+n)&1 == 1 {
		return float64(getKthElement(nums1, nums2, (m+n)/2+1))
	} else {
		r1 := float64(getKthElement(nums1, nums2, (m+n)/2))
		r2 := float64(getKthElement(nums1, nums2, (m+n)/2+1))
		return (r1 + r2) / 2
	}
}

func getKthElement(nums1 []int, nums2 []int, k int) int {
	i1, i2 := 0, 0
	for {
		if i1 == len(nums1) {
			return nums2[i2+k-1]
		}
		if i2 == len(nums2) {
			return nums1[i1+k-1]
		}
		if k == 1 {
			return min(nums1[i1], nums2[i2])
		}
		half := k / 2
		newIndex1 := min(i1+half, len(nums1)) - 1
		newIndex2 := min(i2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - i1 + 1)
			i1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - i2 + 1)
			i2 = newIndex2 + 1
		}
	}
	return 0
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	//fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    //2.0
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) //2.5
	//fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0})) //0
}
