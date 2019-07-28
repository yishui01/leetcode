/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    return helper(root, 0, sum)
}

func helper(root *TreeNode, tmpRes int, sum int) bool {
    if root == nil {
        return false
    }
    nowRes := root.Val + tmpRes 
    if nowRes == sum && root.Left == nil && root.Right == nil{ //注意要叶子节点才算
        return true
    }
    return helper(root.Left,  nowRes, sum) || helper(root.Right,  nowRes, sum)
    
}
