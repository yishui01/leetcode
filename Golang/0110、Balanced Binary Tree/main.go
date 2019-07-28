/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//递归每个节点，比较以该节点为根节点的左右子树的最大深度，看相差是不是超过1，平衡则继续递归该根节点的左右节点，递归时加入 子树如果不平衡 返回 -1 作为剪枝，不用进行下面的递归了
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    rightDep := getDep(root.Right)
    if rightDep == -1 {
        return false
    }
    leftDep := getDep(root.Left)
    if leftDep == -1 {
        return false
    }
    
    if leftDep >= rightDep {
        return (leftDep - rightDep) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
    } else {
        return (rightDep - leftDep) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
    }
    
}

func getDep(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    right := getDep(root.Right)
    if right == -1 {
        return -1
    }
    left := getDep(root.Left)
    if left == -1 {
        return -1
    }
    
    if right >= left {
        if right - left >=2 {
            return -1
        }
        return right + 1
    } else {
        if left - right >= 2 {
            return -1
        }
        return left + 1
    }
    
}
