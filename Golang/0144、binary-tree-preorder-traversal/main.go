/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) []int {
    //迭代流二
    if root == nil {
        return []int{}
    }
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    p := root
    for p != nil || len(stack) > 0{
        if p != nil {
            res = append(res, p.Val)
            stack = append(stack, p)
            p = p.Left
        } else {
            p = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            p = p.Right
        }
    }
    return res
}


// func preorderTraversal(root *TreeNode) []int {
//     //迭代流二
//     if root == nil {
//         return []int{}
//     }
//     res := make([]int, 0)
//     stack := make([]*TreeNode, 0)
//     p := root
//     for p != nil || len(stack) > 0{
//         if p != nil {
//             res = append(res, p.Val)
//             stack = append(stack, p)
//             p = p.Left
//         } else {
//             p = stack[len(stack)-1]
//             stack = stack[:len(stack)-1]
//             p = p.Right
//         }
//     }
//     return res
// }

/**
func preorderTraversal(root *TreeNode) []int {
    //迭代流一
    if root == nil {
        return []int{}
    }
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    for len(stack) != 0 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, top.Val)
        if top.Right != nil{
            stack = append(stack, top.Right)
        }
        if top.Left != nil {
            stack = append(stack, top.Left)
        }
    }
    return res
}
**

/**
func preorderTraversal(root *TreeNode) []int {
    //递归流
    if root == nil {
        return []int{}
    }
    res := make([]int, 0)
    res = append(res, root.Val)
    if root.Left != nil {
        res = append(res, preorderTraversal(root.Left)...)
    }
    if root.Right != nil {
        res = append(res, preorderTraversal(root.Right)...)
    }
    return res
}
**/
