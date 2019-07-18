/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    //？？？一个结点一个结点的比对就行
    
    if p == nil && q == nil {
        return true
    }
    
    if (p == nil && q != nil) || (p != nil && q == nil)  || p.Val != q.Val {
            return false
    }
    
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

