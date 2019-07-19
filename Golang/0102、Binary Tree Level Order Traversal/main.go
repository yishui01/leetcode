/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func levelOrder(root *TreeNode) [][]int {
    
    //本题难点在于你用递归时，左右子树的递归是在不同的函数中，所以，不能同时将当前层级的所有节点中直接加到数组中
    //所以这里加了一个level变量，指明当前递归的层级，以该变量为索引，往res中添加元素，就能保证同一层级的加到同一子数组中
    
    res := make([][]int, 0)
    
    recursion(root, 0, &res)
    
    return res
}


func recursion(node *TreeNode, level int, res *[][]int) {
    if node == nil {
        return 
    }
    
    if len(*res) < level + 1 {
        *res = append(*res, []int{node.Val})
    } else {
        (*res)[level] = append((*res)[level],node.Val )
    }
    
    if node.Left != nil {
        recursion(node.Left, level + 1, res)
    }
    if node.Right != nil {
        recursion(node.Right, level + 1, res)
    }
    
}

//下面的解法是本人是菜鸟级解法，space complexity 达到了O（n）级别，你敢信？ ╮(╯▽╰)╭
// func levelOrder(root *TreeNode) [][]int {
//     res := make([][]int, 0)
    
//     if root == nil {
//         return [][]int{}
//     }
    
//     res = append(res, []int{root.Val})
    
//     recursion([]*TreeNode{root}, &res)
    
//     return res
// }

// func recursion(nodes []*TreeNode,  res *[][]int) {
    
//     if len(nodes) > 0 {
//         next := make([]*TreeNode, 0)
//         tmp := make([]int, 0)
//         for i:=0; i < len(nodes); i++ {
//             if nodes[i].Left !=nil {
//                 tmp = append(tmp, nodes[i].Left.Val)
//                 next = append(next, nodes[i].Left)
//             }
//             if nodes[i].Right !=nil {
//                 tmp = append(tmp, nodes[i].Right.Val)
//                 next = append(next, nodes[i].Right)
//             }
//         }
//         if len(tmp) != 0 {
//             *res = append(*res, tmp)
//             recursion(next, res)
//         }
//     }
    
// }
