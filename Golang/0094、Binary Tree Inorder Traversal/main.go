/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    iteration(root, stack, &res)
    return res
}

func iteration(node *TreeNode, stack []*TreeNode, res *[]int){
    for node != nil || len(stack) > 0 {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        if len(stack) > 0 {
             *res = append(*res, stack[len(stack)-1].Val)
             node = stack[len(stack)-1]
             stack = stack[:(len(stack)-1)]
        }
       
        if node != nil {
            node = node.Right
        }
        
    }
}


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
