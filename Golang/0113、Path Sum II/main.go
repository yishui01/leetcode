/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    //这种要结果集的就是递归回溯了
    res := make([][]int, 0)
    helper(root, sum, 0, []int{}, &res)
    return res
}

func helper(root *TreeNode, sum int, tmpSum int, tmpRes []int, res *[][]int) {
    if root == nil {
        return 
    }
    
    if tmpSum + root.Val == sum && root.Left == nil && root.Right == nil {
        dist := make([]int, len(tmpRes))
        copy(dist, tmpRes)
        dist = append(dist, root.Val)
        *res = append(*res, dist)
        return 
    }
    
    helper(root.Left, sum, tmpSum + root.Val, append(tmpRes, root.Val), res)
    helper(root.Right, sum, tmpSum + root.Val, append(tmpRes, root.Val), res)
    
}


