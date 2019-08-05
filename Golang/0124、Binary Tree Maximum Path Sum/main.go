/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    //这个也是不会做，这个应该算回溯
    if root == nil {
        return 0
    }
    res :=root.Val
    helper(root, &res)
    return res
    
}

func helper(node *TreeNode, max *int) int {
    if node == nil {
        return 0
    }
    
    maxLeft := helper(node.Left, max) 
    maxRight := helper(node.Right, max)
    
    self := node.Val + maxs(0,maxLeft) + maxs(0,maxRight) //这是将当前节点作为根节点进行相加得出来的和
    
    returnRes := node.Val+maxs(0,maxs(maxLeft,maxRight)) //这是将当前节点作为路径节点一部分，准备return给上级用的
    
    if *max < self { //这里只需要将self和max比较就行了，因为self比returnRes多+了一个至少>=0的节点值，所以不会比returnRes小
        *max = self //更新下节点
    }
    
    return returnRes
    
}

func maxs(a,b int)int{
    if a >= b {
        return a
    }
    return b
}
