/*
// Definition for a Node.
class Node {
    public $val;
    public $left;
    public $right;
    public $next;

    @param Integer $val 
    @param Node $left 
    @param Node $right 
    @param Node $next 
    function __construct($val, $left, $right, $next) {
        $this->val = $val;
        $this->left = $left;
        $this->right = $right;
        $this->next = $next;
    }
}
*/
class Solution {

    /**
     * @param Node $root
     * @return Node
     */
    function connect($root) {
        //完全二叉树，左节点指右结点，右结点看parent的Next是不是存在，存在就指向parent的Next的左节点(存在的话)
         if ($root != null) {
            if ($root->left) {
                $root->left->next = $root->right;
            }
             if($root->right){
                 if($root->next){
                    $root->right->next = $root->next->left;
                 } else {
                     $root->righ->next = null;
                 }
             }
             $this->connect($root->left);
             $this->connect($root->right);
        }
        return $root;
    }
    
    
}
