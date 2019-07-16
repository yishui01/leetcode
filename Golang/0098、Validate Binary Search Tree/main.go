/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    //我想问一句，这有什么难度？？好吧还是有个难点，就是需要考虑左子树所有的值要小于所有根节点，右边同样要全部大于根节点
    return recursion(root, 1<<32 - 1,-1 << 32)
}


func recursion(node *TreeNode, max int, min int) bool {
    
    if node == nil {
        return true
    }
   
    if node.Val >=max || node.Val <= min {
        return false
    }
    
    return recursion(node.Left, node.Val, min) &&  recursion(node.Right, max, node.Val)
    
}
