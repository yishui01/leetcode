/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    return helper(root,-1)
}

func helper(root *TreeNode,flag int) int {
     if root == nil{
        return 0
    }
    if root.Left == nil && root.Right == nil && flag == 1 {
        return root.Val
    }
     return helper(root.Left,1) + helper(root.Right,-1)
}
