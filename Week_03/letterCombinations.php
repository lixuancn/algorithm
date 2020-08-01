<?php
//回溯
class Solution {
    private $map = array(
        "2" => "abc",
        "3" => "def",
        "4" => "ghi",
        "5" => "jkl",
        "6" => "mno",
        "7" => "pqrs",
        "8" => "tuv",
        "9" => "wxyz",
    );
    private $ret = array();
    /**
     * @param String $digits
     * @return String[]
     */
    function letterCombinations($digits) {
        if(!$digits){
            return [];
        }
        $this->search('', $digits, 0);
        return $this->ret;
    }

    private function search($ans, $digits, $i)
    {
        if ($i == strlen($digits)){
            $this->ret[] = $ans;
            return;
        }
        for($j=0; $j<strlen($this->map[$digits[$i]]); $j++){
            $char = $this->map[$digits[$i]][$j];
            $ans .= $char;
            $this->search($ans, $digits, $i+1);
            $ans = substr($ans, 0, -1);
        }
    }
}
$ret = (new Solution())->letterCombinations('23');
var_dump($ret);
