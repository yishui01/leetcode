<?php


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

class Solution
{
    //把右边的先连起来，再从左边一直往右边连，这里因为不是完全二叉树，所以需要一直往右边找可连接的节点
    /**
     * @param Node $root
     * @return Node
     */
    function connect ($root)
    {
        if ($root) {
            $pNext = $root->next;
            while ($pNext) {
                if ($pNext->left) {
                    $pNext = $pNext->left;
                    break;
                }
                if ($pNext->right) {
                    $pNext = $pNext->right;
                    break;
                }
                $pNext = $pNext->next;
            }
            if ($root->left) {
                $root->left->next = $root->right ? $root->right : $pNext;
            }
            if ($root->right) {
                $root->right->next = $pNext;
            }

            $this->connect($root->right); //这里要先放right，把右边的树全部横向连接好，不然左子往右边找的时候找不到
            $this->connect($root->left);
        }
        return $root;
    }
}
//
///*
//// Definition for a Node.
//class Node {
//    public $val;
//    public $left;
//    public $right;
//    public $next;
//
//    @param Integer $val
//    @param Node $left
//    @param Node $right
//    @param Node $next
//    function __construct($val, $left, $right, $next) {
//        $this->val = $val;
//        $this->left = $left;
//        $this->right = $right;
//        $this->next = $next;
//    }
//}
//*/
//class Solution {
//
//    /**
//     * @param Node $root
//     * @return Node
//     */
//    //直接将所有节点存起来，按照level分层，然后每层遍历指向右边节点
//    function connect($root) {
//        $allNode = $this->recursion($root, [], 0);
//        foreach($allNode as $v){
//            $count = count($v);
//            foreach($v as $kk => $vv){
//                if($kk < $count-1) {
//                    $v[$kk]->next = $v[$kk+1];
//                }
//            }
//        }
//        return $root;
//    }
//
//    public function recursion($root, $allNode, $level) {
//        if ($root != null){
//            if(isset($allNode[$level])) {
//                $allNode[$level][] = $root;
//            } else {
//                $allNode[$level] = [$root];
//            }
//            $allNode = $this->recursion($root->left, $allNode, $level+1);
//            $allNode = $this->recursion($root->right, $allNode, $level+1);
//        }
//
//        return $allNode;
//    }
//
//
//
////还有一种是一层一层的处理，本层处理的同时构造出下一层的所有节点数组，再将下一层数组传递到下一层
////     function connect($root) {
////         $this->bts([$root]);
////         return $root;
////     }
//
////     function bts($roots)
////     {
////         if (empty($roots)) {
////             return;
////         }
////         $mostLeft = null;
////         $nextRoots = [];
////         foreach ($roots as $root) {
////             if ($root->left) {
////                 if ($mostLeft && ! $mostLeft->next) {
////                     $mostLeft->next = $root->left;
////                 }
////                 $mostLeft = $root->left;
////                 $nextRoots[] = $root->left;
////             }
////             if ($root->right) {
////                 if ($mostLeft && ! $mostLeft->next) {
////                     $mostLeft->next = $root->right;
////                 }
////                 $mostLeft = $root->right;
////                 $nextRoots[] = $root->right;
////             }
////         }
//
////         $this->bts($nextRoots);
////     }
//
//}
