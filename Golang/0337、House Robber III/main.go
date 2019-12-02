/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
   maps := make(map[*TreeNode]int)
  return helper(root,maps)
}

func helper(root *TreeNode,maps map[*TreeNode]int) int {
    if root == nil {
        return 0
    }
    if val,ok := maps[root];ok {
        return val
    }
    money := root.Val
    leftMon,rightMon := 0,0
    if root.Left != nil {
        leftMon += helper(root.Left.Left,maps) + helper(root.Left.Right,maps)
    }
    if root.Right != nil {
        rightMon += helper(root.Right.Left,maps) + helper(root.Right.Right,maps)
    }
    res := max(money+leftMon+rightMon,helper(root.Left,maps)+helper(root.Right,maps))
    maps[root] = res
    return res
    
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
