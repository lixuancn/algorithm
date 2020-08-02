<?php
//贪心
class Solution {

    /**
     * @param Integer[] $commands
     * @param Integer[][] $obstacles
     * @return Integer
     */
    function robotSim($commands, $obstacles) {
        //（dx, dy）就是向上、右、下、左
        $dx = [0, 1, 0, -1];
        $dy = [1, 0, -1, 0];
        //当前左边和方向
        //上右下左分别是0，1，2，3
        $x = $y = $di = 0;
        $obstacleSet = array();
        foreach($obstacles as $obstacle){
            $k = implode('_', $obstacle);
            $obstacleSet[$k] = $obstacle;
        }
        $ans = 0;
        foreach($commands as $cmd){
            if($cmd == -1){
                $di = ($di + 1) % 4;
            }else if($cmd == -2){
                $di = ($di + 3) % 4;
            }else{
                for($i=1; $i<=$cmd; $i++){
                    $nx = $x + $dx[$di];
                    $ny = $y + $dy[$di];
                    $k = $nx . '_' . $ny;
                    //没遇到障碍物
                    if(!isset($obstacleSet[$k])){
                        $x = $nx;
                        $y = $ny;
                        $ans = max($ans, $x*$x+$y*$y);
                    }else{
                        break;
                    }
                }
            }
        }
        return $ans;
    }
}

$ret = (new Solution())->robotSim(
    [4,-1,4,-2,4], [[2,4]]
);
var_dump($ret);
