/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//Morris 遍历了解一下，这个方法主要思路是迭代的同时记录当前的父节点到当前父节点的左子树的最右节点的右孩子上，使得能够一直遍历下去
func inorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    cur := root
    
    for cur != nil {
        if cur.Left == nil {
            res = append(res, cur.Val)
            cur = cur.Right
        } else {
            //找左子树所有节点的最右节点
            pre := cur.Left
            for pre.Right != nil && pre.Right != cur {
                pre = pre.Right
            }
            
            //还没到最右节点的时候，就直接赋值
            if pre.Right == nil {
                pre.Right = cur
                cur = cur.Left
            }
            
            //已经到最右节点了，说明父节点的左子树全部遍历完成，那就可以利用之前保存好的父节点，继续往下遍历
            if pre.Right == cur {
                pre.Right = nil
                res = append(res, cur.Val)
                cur = cur.Right
            }
        }
    }
    
    return res
    
}





//这是用栈来模拟递归，本质思路还是一样的
// func inorderTraversal(root *TreeNode) []int {
//     res := make([]int, 0)
//     stack := make([]*TreeNode, 0)
//     iteration(root, stack, &res)
//     return res
// }

// func iteration(node *TreeNode, stack []*TreeNode, res *[]int){
//     for node != nil || len(stack) > 0 {
//         for node != nil {
//             stack = append(stack, node)
//             node = node.Left
//         }
//         if len(stack) > 0 {
//              *res = append(*res, stack[len(stack)-1].Val)
//              node = stack[len(stack)-1]
//              stack = stack[:(len(stack)-1)]
//         }
       
//         if node != nil {
//             node = node.Right
//         }
        
//     }
// }


 //先用递归试试
// func inorderTraversal(root *TreeNode) []int {
//     res := make([]int, 0)
//     recursion(root, &res)
//     return res
// }

// func recursion(node *TreeNode, res *[]int) {
//     if node == nil {
//         return 
//     }
    
//     if node.Left != nil {
//         recursion(node.Left, res)
//     }
//     *res = append(*res, node.Val)
//     if node.Right != nil {
//         recursion(node.Right, res)
//     }
// }