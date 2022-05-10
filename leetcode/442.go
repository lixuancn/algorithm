package main

//442. 数组中重复的数据

func findDuplicates(nums []int) []int {
    // return findDuplicates_tag(nums) //标记法
    return findDuplicates_swap(nums) //交换法
}

//标记法，数字x，就把x-1的下标变成负数
func findDuplicates_tag(nums []int) []int {
    result := make([]int, 0)
    for _, x := range nums{
        if x < 0{
            x = -x
        }
        if nums[x-1] < 0{
            result = append(result, x)
        }else{
            nums[x-1] = -nums[x-1]    
        }
    }
    return result
}

//交换法，把x放在下标x-1的地方，最后循环一次，如果x!=i+1，则说明x没放到规定的地方，则说明x的位置已经被占了，那么x就出现了2次
func findDuplicates_swap(nums []int) []int {
    result := make([]int, 0)
    for i := range nums{
        for nums[i] != nums[nums[i]-1]{
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }
    for i := range nums{
        if nums[i] != i+1{
            result = append(result, nums[i])
        }
    }
    return result
}
