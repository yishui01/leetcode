/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    //我知道了，后序遍历的要点就是根节点在最后，知道这一个，就和前一题一样的啦,分治
    inLen := len(inorder)
    postLen := len(postorder)
    
    if inLen == 0 || inLen != postLen {
        return nil
    }
    root := new(TreeNode)
    root.Val = postorder[postLen-1]
    
    for i:=0; i < inLen; i++ {
        //在中序中找到根节点，将其分为左右子树
        if inorder[i] == root.Val {
            left := i  //左子树的元素个数
            right := inLen - i - 1 //右子树元素的个数
            if left > 0 {
                root.Left = buildTree(inorder[:left], postorder[:left])
            }
            if right > 0 {
                root.Right = buildTree(inorder[left+1:], postorder[left:inLen-1])
            }
            break
        }
    }
    
    return root
}
