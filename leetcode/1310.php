<?php
//有一个正整数数组 arr，现给你一个对应的查询数组 queries，其中 queries[i] = [Li, Ri]。
//
//对于每个查询 i，请你计算从 Li 到 Ri 的 XOR 值（即 arr[Li] xor arr[Li+1] xor ... xor arr[Ri]）作为本次查询的结果。
//
//并返回一个包含给定查询 queries 所有结果的数组。

//暴力遍历会超时，自然而然想到缓存结果，缓存是个技巧
//缓存的是cacheList[0] = arr[0]
//cacheList[1] = cacheList[0]^arr[1]
//cacheList[2] = cacheList[1]^arr[2]
//cacheList[N] = cacheList[N-1]^arr[N]
//接下是个数学证明，借助辅助变量
//Q(0, right) = cacheList[right]
//Q(left,right) = arr[left]^...^arr[right]
// = x ^ x ^ (arr[left]^...^arr[right])
// = (arr[0]^...^arr[left-1]) ^ (arr[0]^...^arr[left-1]) ^ (arr[left]^...^arr[right])
// = (arr[0]^...^arr[left-1]) ^ (arr[0]^...^arr[left-1]^arr[left]^...^arr[right])
// = (arr[0]^...^arr[left-1]) ^ cacheList[right]
// = cacheList[left-1] ^ cacheList[right]

class Solution {
    /**
     * @param Integer[] $arr
     * @param Integer[][] $queries
     * @return Integer[]
     */
    function xorQueries($arr, $queries) {
        $ret = array();
        $cacheList = array(-1=>0);
        foreach($arr as $k=>$a){
            $cacheList[$k] = $cacheList[$k-1] ^ $a;
        }
        foreach($queries as $query){
            $r = $cacheList[$query[0]-1] ^ $cacheList[$query[1]];
            $ret[] = $r;
        }
        return $ret;
    }

     //O(mn) 超限
    function xorQueries_loop($arr, $queries) {
        $ret = array();
        foreach($queries as $query){
            $r = 0;
            for($i=$query[0]; $i<=$query[1]; $i++){
                $r ^= $arr[$i];
            }
            $ret[] = $r;
        }
        return $ret;
    }
}

$result = (new Solution())->xorQueries([1,3,4,8], [[0,1],[1,2],[0,3],[3,3]]);
var_dump($result);
//[2,7,14,8]
//数组中元素的二进制表示形式是：
//1 = 0001
//3 = 0011
//4 = 0100
//8 = 1000
//查询的 XOR 值为：
//[0,1] = 1 xor 3 = 2
//[1,2] = 3 xor 4 = 7
//[0,3] = 1 xor 3 xor 4 xor 8 = 14
//[3,3] = 8
