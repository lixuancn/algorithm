<?php
//丑数。堆的方式

class Solution {
    /**
     * @param Integer $n
     * @return Integer
     */
    function nthUglyNumber($n) {
        $heap = new Heap();
        $seen = array(1=>0);
        $primes = array(2, 3, 5);
        $ret = 1;
        for($i=1; $i<$n; $i++){
            foreach($primes as $prime){
                $val = $prime * $ret;
                if(!isset($seen[$val])){
                    $heap->insert($val, $val);
                    $seen[$val] = 0;
                }
            }
            $ret = $heap->extract();
        }
        return $ret;
    }
}

class Heap extends \SplPriorityQueue{
    public function compare($priority1, $priority2)
    {
        return $priority2 - $priority1;
    }
}

$r = (new Solution())->nthUglyNumber(11);
var_dump($r);
