/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    //这题有个坑就是要找到叶子节点，从叶子节点到root才算最短路径，一个节点有左孩子但没有右孩子，那么它的右孩子到root的路径是不作数的
    //那么就要避免找本身为nil的节点，而是要找自己有值，却没有孩子的节点
    if root == nil {
        return 0
    }
    return getDep(root)
}

func getDep(root *TreeNode) int {
    if root == nil {
        return -1 //这里本身就是nil节点，不合法，返回 -1 用于剪枝
    }
    if root.Left == nil && root.Right == nil { //这里找到了叶子节点，开始递归
        return 1
    }
    left:=getDep(root.Left)
    right := getDep(root.Right)
   
    if left == -1 && right == -1 {
        return 1
    } else if left == -1 {
        return right + 1
    } else if right == -1{
        return left + 1
    } else {
        if left >= right {
            return right + 1
        } else {
            return left + 1
        }
    }
    
}


//拜读打来解法，递归写的就是优雅
// func minDepth(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     if root.Left == nil { //直接避免了递归本身为nil的节点
//         return 1+minDepth(root.Right)
//     }
//     if root.Right == nil {
//         return 1+minDepth(root.Left)
//     }
    
//     return 1+min(minDepth(root.Right),minDepth(root.Left))

// }

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
    
//     return b
// }

