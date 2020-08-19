<?php
//查并集

class UnionFind{
    public $parent = array();
    private $rank = array();
    private $count = 0;
    public function __construct($n){
        $this->count = $n;
        for($i=0; $i<$n; $i++){
            $this->parent[$i] = $i;
            $this->rank[$i] = 1;
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
        if($this->rank[$pRoot] > $this->rank[$qRoot]){
            $this->parent[$qRoot] = $pRoot;
        }else if($this->rank[$pRoot] < $this->rank[$qRoot]){
            $this->parent[$pRoot] = $qRoot;
        }else{
            $this->parent[$qRoot] = $pRoot;
            $this->rank[$pRoot]++;
        }
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
    private $m = 0;
    private $n = 0;

    /**
     * @param String[][] $grid
     * @return Integer
     */
    function numIslands($grid) {
        if(!$grid){
            return 0;
        }
        $this->m = count($grid);
        $this->n = count($grid[0]);
        // 多开一个空间，把水域 "0" 都归到这个虚拟的根上
        $dummyNode = $this->m * $this->n;
        $uf = new UnionFind($this->m * $this->n + 1);
        for($i=0; $i<$this->m; $i++){
            for($j=0; $j<$this->n; $j++) {
                # 如果是水域，都连到那个虚拟的空间去
                if($grid[$i][$j] == 0){
                    $uf->union($this->getIndex($i, $j), $dummyNode);
                }else if($grid[$i][$j] == 1){
                    foreach(array(array(1, 0), array(0, 1)) as $pos){
                        $newI = $i + $pos[0];
                        $newJ = $j + $pos[1];
                        if($newI < $this->m && $newJ < $this->n && $grid[$newI][$newJ] == 1){
                            $uf->union($this->getIndex($i, $j), $this->getIndex($newI, $newJ));
                        }
                    }
                }
            }
        }
        return $uf->getCount() - 1;
    }

    private function getIndex($x, $y){
        return $x * $this->n + $y;
    }
}

$obj = new Solution();
$r = $obj->numIslands(
    [
        ['1','1','0','0','0'],
        ['1','1','0','0','0'],
        ['0','0','1','0','0'],
        ['0','0','0','1','1']
    ]
);
var_dump($r);
