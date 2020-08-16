<?php
//排序+循环，从数量最多的任务开始，n为一组，每组都是数量最多的任务开头
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
        $ans = 0;
        while($map[25]){
            $i = 0;
            while ($i <= $n){
                if($map[25] == 0){
                    break;
                }
                if($i < 26 && $map[25-$i] > 0){
                    $map[25-$i]--;
                }
                $ans++;
                $i++;
            }
            sort($map);
        }
        return $ans;
    }
}
$ret = (new Solution())->leastInterval(
    ["A","A","A","B","B","B"], 2
);
var_dump($ret);
