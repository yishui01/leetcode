/*
// Definition for a Node.
class Node {
    public $val;
    public $next;
    public $random;

    @param Integer $val 
    @param Node $next 
    @param Node $random 
    function __construct($val, $next, $random) {
        $this->val = $val;
        $this->next = $next;
        $this->random = $random;
    }
}
*/
class Solution {

    /**
     * @param Node $head
     * @return Node
     */
    function copyRandomList($head) {
        $array = [];
        return $this->helper($head,$array);
    }
    
    function helper($node, &$array){
        if (!$node){
            return null;
        }
        $key = 'A-'.$node->val;
        if (isset($array[$key])) {
            $clone = $array[$key];
        } else {
            $clone = clone $node;
            $array[$key] = $clone;
        }
       
        //遍历next
        if ($clone->next){
            $clone->next = $this->helper($clone->next, $array);
        }
        //遍历random
        if ($clone->random){
            //这里由于next都已经跑过一遍了，那也就是说所有node都已经复制好了，此时random指向哪个就直接拿来用就行
            $keyR = 'A-'.$node->random->val;
            $clone->random = $array[$keyR];
        }
        
        return $clone;
    }
}
