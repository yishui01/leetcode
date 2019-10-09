/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    //二叉树中序遍历是升序，直接利用这个性质即可，遍历到第k个的时候，就是结果,这里使用栈+迭代
    stack := make([]*TreeNode,0)
    count := 0
    
    p := root 
    
    for len(stack) > 0 || p != nil {
        for p != nil {
            stack = append(stack, p)
            p = p.Left
        }
        p = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        count++
        if count == k {
            return p.Val
        }
        p = p.Right
    }
    return -1 
}



// func kthSmallest(root *TreeNode, k int) int {
//     //二叉树中序遍历是升序，直接利用这个性质即可，遍历到第k个的时候，就是结果
//     res := make([]int,0)
//     dfs(root,k,&res)
//     return res[len(res)-1]  
// }

// func dfs(node *TreeNode,k int,res *[]int) {
//     if node == nil {
//         return 
//     }
//     if node.Left != nil {
//         dfs(node.Left,k,res)
//     }
    
//     if k == len(*res) {
//         return 
//     }
    
//     *res = append(*res, node.Val)
     
//     if k == len(*res) {
//         return 
//     }
    
//     if node.Right != nil {
//         dfs(node.Right,k,res)
//     }
    
// }
