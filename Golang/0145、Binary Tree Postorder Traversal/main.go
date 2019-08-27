/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    //后序遍历，栈的方法，这里是要将全部节点都要压入栈中，使用一个head辅助指针，记录之前被加入到res中的node
    //加入到res可以有三个条件，要么是叶子节点，因为压栈顺序已经摆好了，所以叶子节点直接加入到res中即可
    //然后弹出下一个stack的顶部节点，顶部的节点要确定是否为可加入res的根节点，什么叫可加入的根节点？就是他的左右孩子都已经加入res了
    //那么他也可以加入了，那之前的head指针就起作用了，如果head == 当前弹出节点的左孩子或者右孩子，那当前节点也可以加入了
    if root == nil {
        return []int{}
    }
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    head := root 
    for len(stack) > 0{
        node := stack[len(stack)-1]
        if (node.Left == nil && node.Right == nil) || node.Left == head || node.Right == head {
            res = append(res, node.Val)
            head = node
            stack = stack[:len(stack)-1]
        } else {
            if node.Right != nil {
              stack = append(stack, node.Right)
            }
            if node.Left != nil  {
              stack = append(stack, node.Left)
            }
        }
    }
    
    return res

}

/***
func postorderTraversal(root *TreeNode) []int {
    //后序遍历，先递归试试
    if root == nil {
        return []int{}
    }
    res := make([]int, 0)
    if root.Left != nil{
        res = append(res, postorderTraversal(root.Left)...)
    }
    if root.Right != nil {
        res = append(res, postorderTraversal(root.Right)...)
    }
    
   res = append(res, root.Val)
    
    return res

}
***/

