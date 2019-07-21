/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//本菜鸟的递归还是有差距
func maxDepth(root *TreeNode) int {
    var res int
    
    recursion(root, 0,&res)
    
    return res
}

func recursion(node *TreeNode, level int , res *int) {
    if node == nil {
        return 
    }
    if level >= *res {
        *res = *res + 1
    }
    if node.Left != nil {
        recursion(node.Left, level + 1 ,res)
    }
     if node.Right != nil {
        recursion(node.Right, level + 1 ,res)
    }
}

// //大佬的递归就是这么6
// func maxDepth(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
    
//     lr := maxDepth(root.Left)
//     rr := maxDepth(root.Right)
//     if lr > rr {
//         return 1 + lr
//     } else {
//         return 1 + rr
//     }
// }


