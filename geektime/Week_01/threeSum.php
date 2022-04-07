<?php

/**
 * 先排序，再用双指针。
 * Class Solution
 */
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function threeSum($nums) {
        $len = count($nums);
        if($len < 3){
            return array();
        }
        sort($nums);
        $res = array();
        foreach($nums as $i => $num){
            if($num > 0){
                break;
            }
            //过滤重复
            if(isset($nums[$i-1]) && $num == $nums[$i-1]){
                continue;
            }
            $l = $i + 1;
            $r = $len - 1;
            while($l < $r){
                //若和小于 00，说明 nums[L]nums[L] 太小，LL 右移
                if($num + $nums[$l] + $nums[$r] < 0){
                    $l++;
                //若和大于 0，说明 nums[R]nums[R] 太大，RR 左移
                }else if($num + $nums[$l] + $nums[$r] > 0){
                    $r--;
                }else{
                    $res[] = array($num, $nums[$l], $nums[$r]);
                    //避开重复
                    while($l < $r && $nums[$l] == $nums[$l+1]){
                        $l++;
                    }
                    while($l < $r && $nums[$r] == $nums[$r+1]){
                        $r--;
                    }
                    $l++;
                    $r--;
                }
            }
        }
        return $res;
    }
}
$obj = new Solution();
$nums = $obj->threeSum([-1, 0, 1, 2, -1, -4]);
print_r($nums);
