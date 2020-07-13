<?php
class Solution {

    /**
     * @param String[] $strs
     * @return String[][]
     */
    function groupAnagrams($strs) {
        //答案是抄的
        //用质数来做map的key，除了1和它自身外，不能被其他自然数整除的数叫做质数。
        //根据答案，我觉得：任意N个质数相乘，值都不相同是质数的特性。
        $resArr = [];
        // 将 26 个字母映射为 素数，求积可得唯一 key，相当于一个无冲突的 hash function
        $prime = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103];
        foreach ($strs as $str) {
            $strlen = 1;
            for ($i = 0; $i < strlen($str); $i++) {
                $strlen *= $prime[ord($str[$i]) - 97];
            }
            $resArr[$strlen][] = $str;
        }
        return array_values($resArr);
    }
}

$r = (new Solution())->groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]);
var_dump($r);
