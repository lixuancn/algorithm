<?php
//查并集
class UnionFind{
    private $parent = array();
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

    public function isConnect($p, $q){
        return $this->find($p) == $this->find($q);
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
}
