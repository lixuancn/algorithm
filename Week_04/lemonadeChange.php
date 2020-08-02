<?php
//贪心
class Solution {

    /**
     * @param Integer[] $bills
     * @return Boolean
     */
    function lemonadeChange($bills) {
        if(!$bills){
            return false;
        }
        if($bills[0] != 5 || $bills[1] == 20){
            return false;
        }
        $cash = array(20=>0, 10=>0, 5=>0);
        foreach($bills as $bill){
            $cash[$bill]++;
            $sub = $bill - 5;
            if($sub == 0){
                continue;
            }
            $tmp = $cash;
            foreach($cash as $c => $n){
                if($n == 0){
                    continue;
                }
                if($sub >= $c){
                    $num = intval($sub / $c);
                    if($num > $n){
                        continue;
                    }
                    $sub = $sub - $c * $num;
                    $tmp[$c] = $tmp[$c] - $num;
                }
            }
            if($sub != 0){
                return false;
            }
            $cash = $tmp;
        }
        return true;
    }

}

$ret = (new Solution())->lemonadeChange(
    [5,5,5,10,20]
);
var_dump($ret);
