<?php
//双向DFS
class Solution {

    /**
     * @param String $start
     * @param String $end
     * @param String[] $bank
     * @return Integer
     */
    function minMutation($start, $end, $bank) {
        if(!$start || !$end || !$bank || !in_array($end, $bank)){
            return -1;
        }
        if($start == $end){
            return 0;
        }
        $visited = array();
        $beginSet = array($start=>1);
        $endSet = array($end=>1);
        $wordSet = array();
        foreach($bank as $word){
            $wordSet[$word] = 1;
        }
        $step = 0;
        while($beginSet && $endSet){
            if(count($beginSet) > count($endSet)){
                $tmp = $beginSet;
                $beginSet = $endSet;
                $endSet = $tmp;
            }
            $tmpSet = array();
            foreach($beginSet as $word => $v){
                for($i=0; $i<strlen($word); $i++){
                    foreach(array('A', 'C', 'G', 'T') as $c){
                        $w = $word;
                        $w[$i] = $c;
                        if(isset($endSet[$w])){
                            return $step+1;
                        }
                        if(!isset($visited[$w]) && isset($wordSet[$w])){
                            $tmpSet[$w] = 1;
                            $visited[$w] = 1;
                        }
                    }
                }
            }
            $step++;
            $beginSet = $tmpSet;
        }
        return -1;
    }
}

$obj = new Solution();
$r = $obj->minMutation('AACCGGTT', 'AAACGGTA', ["AACCGGTA", "AACCGCTA", "AAACGGTA"]);
var_dump($r);
