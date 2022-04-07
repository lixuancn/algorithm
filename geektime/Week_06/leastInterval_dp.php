<?php
//DP
class Solution {

    /**
     * @param String[] $tasks
     * @param Integer $n
     * @return Integer
     */
    function leastInterval($tasks, $n) {
        $map = array_fill(0, 26, 0);
        foreach($tasks as $task){
            $map[ord($task)-ord('A')]++;
        }
        sort($map);
        //数量最多的任务有多少个间隔，
        $maxVal = $map[25] - 1;
        //这个间隔可以插入其他的任务，个数是间隔*n
        $idleSlots = $maxVal * $n;
        for($i=24; $i>=0&&$map[$i]>0; $i--){
            //空闲的间隔可以安插别的任务，能安插的数量是min(当前任务数量，间隔个数)
            $idleSlots = $idleSlots - min($map[$i], $maxVal);
        }
        return $idleSlots > 0 ? $idleSlots + count($tasks) : count($tasks);
    }
}
$ret = (new Solution())->leastInterval(
    ["A","A","A","B","B","B"], 2
);
var_dump($ret);
