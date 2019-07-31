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
    //直接将所有节点存起来，按照level分层，然后每层遍历指向右边节点
    function connect($root) {
        $allNode = $this->recursion($root, [], 0);
        foreach($allNode as $v){
            $count = count($v);
            foreach($v as $kk => $vv){
                if($kk < $count-1) {
                    $v[$kk]->next = $v[$kk+1];
                }
            }
        }
        return $root;
    }
    
    public function recursion($root, $allNode, $level) {
        if ($root != null){
            if(isset($allNode[$level])) {
                $allNode[$level][] = $root;
            } else {
                $allNode[$level] = [$root];
            }
            $allNode = $this->recursion($root->left, $allNode, $level+1);
            $allNode = $this->recursion($root->right, $allNode, $level+1);
        }
        
        return $allNode;
    }
    
    
    
//还有一种是一层一层的处理，本层处理的同时构造出下一层的所有节点数组，再将下一层数组传递到下一层
//     function connect($root) {
//         $this->bts([$root]);
//         return $root;
//     }
    
//     function bts($roots)
//     {
//         if (empty($roots)) {
//             return;
//         }
//         $mostLeft = null;
//         $nextRoots = [];
//         foreach ($roots as $root) {
//             if ($root->left) {
//                 if ($mostLeft && ! $mostLeft->next) {
//                     $mostLeft->next = $root->left;
//                 }
//                 $mostLeft = $root->left;
//                 $nextRoots[] = $root->left;
//             }
//             if ($root->right) {
//                 if ($mostLeft && ! $mostLeft->next) {
//                     $mostLeft->next = $root->right;
//                 }
//                 $mostLeft = $root->right;
//                 $nextRoots[] = $root->right;
//             }
//         }
        
//         $this->bts($nextRoots);
//     }
    
}
