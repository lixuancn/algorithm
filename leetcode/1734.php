<?php
//给你一个整数数组 perm ，它是前 n 个正整数的排列，且 n 是个 奇数 。
//它被加密成另一个长度为 n - 1 的整数数组 encoded ，满足 encoded[i] = perm[i] XOR perm[i + 1] 。比方说，如果 perm = [1,3,2] ，那么 encoded = [2,1] 。
//比方说，如果 perm = [1,3,2] ，那么 encoded = [2,1] 。
//
//给你 encoded 数组，请你返回原始数组 perm 。题目保证答案存在且唯一。

//核心思路是求perm[0]
//连续正整数，所以可以1^2^3^...^n 即= perm[0]^perm[1]^perm[2]^....^perm[n]
//perm[0] = (perm[0]^perm[1]^perm[2]^....^perm[n]) ^ (perm[1]^perm[2]^....^perm[n])
//perm[1]^perm[2]^....^perm[n] 的值其实就是 encoded的奇数下标求XOR
class Solution {

    /**
     * @param Integer[] $encoded
     * @return Integer[]
     */
    function decode($encoded) {
        $total = 0;
        for($i=1; $i<=count($encoded)+1; $i++){
            $total ^= $i;
        }
        $odd = 0;
        for($i=1; $i<count($encoded); $i+=2){
            $odd ^= $encoded[$i];
        }
        $ret = array();
        $ret[0] = $odd ^ $total;
        foreach ($encoded as $k=>$v){
            $ret[$k+1] = $v ^ $ret[$k];
        }
        return $ret;
    }
}

$encoded = [2,1];
$result = (new Solution())->decode($encoded);
var_dump($result);
