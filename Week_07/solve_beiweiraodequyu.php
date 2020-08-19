<?php
//查并集

class UnionFind{
    public $parent = array();
    private $count = 0;
    public function __construct($n){
        $this->count = $n;
        for($i=0; $i<$n; $i++){
            $this->parent[$i] = $i;
        }
    }

    //查找根节点
    public function find($p){
        while($p != $this->parent[$p]){
            $this->parent[$p] = $this->parent[$this->parent[$p]];
            $p = $this->parent[$p];
        }
        return $p;
    }

    public function union($p, $q){
        $pRoot = $this->find($p);
        $qRoot = $this->find($q);
        if($pRoot == $qRoot){
            return;
        }
        $this->parent[$qRoot] = $pRoot;
//        $this->parent[$pRoot] = $qRoot;
        $this->count--;
    }

    public function isConnect($p, $q){
        return $this->find($p) == $this->find($q);
    }

    public function getCount(){
        return $this->count;
    }
}

class Solution {
    /**
     * @param String[][] $board
     * @return NULL
     */
    function solve(&$board) {
        if(!$board){
            return 0;
        }
        $m = count($board);
        $n = count($board[0]);
        //O都放到虚拟节点上
        $dummyNode = $m * $n;
        $uf = new UnionFind($m * $n + 1);
        for($i=0; $i<$m; $i++){
            for($j=0; $j<$n; $j++){
                if($board[$i][$j] != 'O') {
                    continue;
                }
                //边界上的O
                if($i==0 || $j==0 || $i==$m-1 || $j==$n-1){
                    $uf->union($this->getIndex($i, $j, $n), $dummyNode);
                }else{
                    foreach(array(array(-1, 0), array(0, -1), array(1, 0), array(0, 1)) as $pos) {
                        $newI = $i + $pos[0];
                        $newJ = $j + $pos[1];
                        if ($newI >= 0 && $newI <= $m-1 && $newJ >= 0 && $newJ <= $n-1 && $board[$newI][$newJ] == 'O') {
                            $uf->union($this->getIndex($i, $j, $n), $this->getIndex($newI, $newJ, $n));
                        }
                    }
                }
            }
        }
        print_r($uf->parent);
        for($i=0; $i<$m; $i++){
            for($j=0; $j<$n; $j++){
                if($board[$i][$j] == 'O' && !$uf->isConnect($this->getIndex($i, $j, $n), $dummyNode)){
                    $board[$i][$j] = 'X';
                }
            }
        }
    }

    private function getIndex($x, $y, $n){
        return $x * $n + $y;
    }
}

$obj = new Solution();
$a = [["O","O","O"],["O","O","O"],["O","O","O"]];
$r = $obj->solve(
    $a
);
var_dump($a);
