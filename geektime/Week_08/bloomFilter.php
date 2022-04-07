<?php
class SimpleHash{
    private $cap;
    private $seed;
    public function __construct($cap, $seed)
    {
        $this->cap = $cap;
        $this->seed = $seed;
    }

    public function hash($value){
        $value = strval($value);
        $result = 0;
        $len = strlen($value);
        for($i=0; $i<$len; $i++){
            $result = $this->seed * $result + $value[$i];
        }
        return ($this->cap) & $result;
    }
}

class BloomFilter{
    const DEFAULT_SIZE = 100;
    private $seeds = array(5, 7, 11, 13, 31, 37, 61);
    private $bits = array();
    private $funcs = array();
    public function __construct()
    {
        $this->bits = array_fill(0, self::DEFAULT_SIZE, false);
        foreach($this->seeds as $seed){
            $this->funcs[] = new SimpleHash(self::DEFAULT_SIZE, $seed);
        }
    }

    public function add($value) {
        foreach($this->funcs as $func){
            $this->bits[$func->hash($value)] = true;
        }
    }

    public function contains($value) {
        if (!$value) {
            return false;
        }
        foreach($this->funcs as $func){
            $result = $this->bits[$func->hash($value)];
            if(!$result){
                return false;
            }
        }
        return true;
    }
}

$obj = new BloomFilter();
$obj->add(10);
$obj->add(15);
$obj->add(20);
var_dump($obj->contains(10));
var_dump($obj->contains(15));
var_dump($obj->contains(20));
var_dump($obj->contains(30));
var_dump($obj->contains(40));
