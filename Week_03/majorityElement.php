<?php
//方法一：暴力统计
//方法二：排序后取中间值
//方法三：随机，先随机一个数，然后统计这个数的次数看是不是众数，不是就继续随机，最坏是O(无穷)
//方法四：分治，如下：

class Solution {
    private $ret = array();
    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function majorityElement($nums) {
        return $this->rec($nums, 0, count($nums)-1);
    }

    private function rec($nums, $lo, $hi){
        if($lo == $hi){
            return $nums[$lo];
        }
        $mid = intval(($hi - $lo) / 2) + $lo;
        $left = $this->rec($nums, $lo, $mid);
        $right = $this->rec($nums, $mid + 1, $hi);
        if($left == $right){
            return $left;
        }
        $leftCount = $this->count($nums, $left, $lo, $mid);
        $rightCount = $this->count($nums, $right, $mid + 1, $hi);
        return $leftCount > $rightCount ? $left : $right;
    }

    private function count($nums, $num, $lo, $hi){
        $count = 0;
        for($i=$lo; $i<=$hi; $i++){
            if($nums[$i] == $num){
                $count++;
            }
        }
        return $count;
    }
}

$ret = (new Solution())->majorityElement([1,3,3]);
var_dump($ret);
