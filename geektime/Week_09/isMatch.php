<?php
//思路：https://leetcode-cn.com/problems/regular-expression-matching/solution/ji-yu-guan-fang-ti-jie-gen-xiang-xi-de-jiang-jie-b/
class Solution {
    private $mem = array();
    /**
     * @param String $s
     * @param String $p
     * @return Boolean
     */
    function isMatch($s, $p) {
        return $this->dp($s, $p, 0, 0);
    }

    function dp($s, $p, $i, $j){
        if($j == strlen($p)){
            return $i == strlen($s);
        }
        if($i == strlen($s)){
            // 如果能匹配空串，一定是字符和 * 成对儿出现
            if((strlen($p) - $j) % 2 == 1){
                return false;
            }
            // 检查是否为 x*y*z* 这种形式
            for(; $j+1 < strlen($p); $j+=2){
                if($p[$j+1] != '*'){
                    return false;
                }
            }
            return true;
        }
        $key = $i . '_' . $j;
        if(isset($this->mem[$key])){
            return $this->mem[$key];
        }
        if($s[$i] == $p[$j] || $p[$j] == '.'){
            if(isset($p[$j+1]) && $p[$j+1] == '*'){
                //通配符匹配 0 次或多次
                $res = $this->dp($s, $p, $i, $j+2) || $this->dp($s, $p, $i+1, $j);
            }else{
                $res = $this->dp($s, $p, $i+1, $j+1);
            }
        }else{
            if(isset($p[$j+1]) && $p[$j+1] == '*'){
                $res = $this->dp($s, $p, $i, $j+2);
            }else{
                $res = false;
            }
        }
        $this->mem[$key] = $res;
        return $res;
    }
}
$obj = new Solution();
$r = $obj->isMatch(
    "ab", ".*c"
);
var_dump($r);
