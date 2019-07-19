/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    return recursion(root.Left, root.Right)
    
}

func recursion(leftNode *TreeNode, rightNode *TreeNode) bool {
    if leftNode == nil && rightNode == nil {
        return true
    }
   
    if leftNode == nil || rightNode == nil {
        return false
    }
    if leftNode.Val == rightNode.Val {
         return recursion(leftNode.Left, rightNode.Right) && recursion(leftNode.Right, rightNode.Left)
    }
   
    return false
}
