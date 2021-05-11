<?php
//密成另一个长度为 n - 1 的整数数组 encoded ，满足 encoded[i] = perm[i] XOR perm[i + 1] 。
//比方说，如果 perm = [1,3,2] ，那么 encoded = [2,1] 。
//
//给你 encoded 数组，请你返回原始数组 perm 。题目保证答案存在且唯一。
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
