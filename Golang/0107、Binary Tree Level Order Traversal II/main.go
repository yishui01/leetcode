/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//思路：先按顺序再反序
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil //nil是可以作为[][]int的返回值的
    }
    res := make([][]int, 0)
    recursion(root, 0, &res)
    
    lens  := len(res)
    for i:=0; i < lens/2; i++ {
        res[i],res[lens-i-1] = res[lens-i-1],res[i]
    }
    
    return res
}

func recursion(node *TreeNode, level int, res *[][]int) {
    if node == nil {
        return 
    }
    if level <= len(*res)-1 {
        (*res)[level] = append((*res)[level], node.Val)
    } else {
        *res = append(*res, []int{node.Val})
    }
    if node.Left != nil {
        recursion(node.Left, level + 1, res)
    }
    if node.Right != nil {
        recursion(node.Right, level + 1 ,res)
    }
}

//看了答案中大佬的解法，是逐步调整res的值 res = append(tmp, res...) 以此达到反序，这种我觉得效率可能还没我这个高
