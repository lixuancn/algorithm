<?php
//还是广度优先

class Solution {

    /**
     * @param String $start
     * @param String $end
     * @param String[] $bank
     * @return Integer
     */
    function minMutation($start, $end, $bank) {
        $possible = array('A', 'C', 'G', 'T');
        $queue = array();
        //字符串，和当前第几步
        array_push($queue, array($start, 0));
        while($queue){
            list($word, $step) = array_shift($queue);
            if($word == $end){
                return $step;
            }
            for($i=0; $i<strlen($word); $i++){
                foreach($possible as $p){
                    //替换一个字符
                    $tmp = $word;
                    $tmp[$i] = $p;
                    //检测合法性
                    foreach($bank as $k=>$b){
                        if($b == $tmp){
                            //避免重复
                            unset($bank[$k]);
                            array_push($queue, array($tmp, $step+1));
                        }
                    }
                }
            }
        }
        return -1;
    }
}
