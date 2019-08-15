/*
// Definition for a Node.
class Node {
    public $val;
    public $neighbors;

    @param Integer $val 
    @param list<Node> $neighbors 
    function __construct($val, $neighbors) {
        $this->val = $val;
        $this->neighbors = $neighbors;
    }
}
*/
class Solution {
//哎哟我的哥啊，第一次做这种题，一度怀疑人生，话说这题没有golang版本的，只能用php了
//思路：dfs
//要注意的点：
//1、无向图会成环，要避免无限循环
//2、最重要的一点，递归的时候要想清楚，第一个是结束递归的条件，第二个是递归函数中起作用的点，真正发力的代码只能有一个
    //比如 dfs最开始就有个 if (isset($array[$k])) { return $array[$k];} ，这里已经判断是否重复了，那么你在下面的for循环中
    //里面的else分支遍历到新的node的时候，不要在else中直接clone，要把他传到下一个dfs，否则你在else中clone了，然后把clone的对象存在array中
    //那么你接下来打算clone这个已经clone的对象中的neighbors，然后把这个对象传到下一个dfs，然后下一个dfs查找array发现本对象已经clone过了，
    //就直接return了，并不会循环这个对象中的neighbors......这就尴尬了，要想清楚再写
//3、本题说没有重复的数字，这还好，可以直接用关联数组进行查找已经生成过的node，有的话直接取出
    
    /**
     * @param Node $node
     * @return Node
     */
    function cloneGraph($node) {
        $arr = [];
        return $this->dfs($node, $arr);
    }
    
    function dfs($node, &$array) {
        if(!$node){
            return ;
        }
        $key = 'A-'.($node->val);
        if (isset($array[$k])){
            return $array[$k];
        }
        $clone = clone $node;
        $array[$key] = $clone;
        $count = count($clone->neighbors);
        for($i=0; $i < $count; $i++) {
            $key = 'A-'.($clone->neighbors[$i]->val);
            if(isset($array[$key])){
                $clone->neighbors[$i] = $array[$key];
            } else {
                $clone->neighbors[$i] = $this->dfs($clone->neighbors[$i], $array);
            }
        }
            
        return $clone;
    }
}
