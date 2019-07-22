/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//本题有两个点，先序的第一个元素是根节 =》 在中序中找到这个根节点，将其分为左右两部分
//还有一个点，中序中找到左边部分的时候，可以通过左边的部分或者右边部分元素的个数，同样找到先序中对应的左子树和右子树的部分
//再将两者的左子树部分和右子树部分作为先序和中序传递到下一层，递归+分治，6到极致
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(preorder) != len(inorder) {
        return nil
    }
    
    r := preorder[0]
    root := &TreeNode{
        Val: r,
    }
    
    for i, v := range inorder {
        if v == r {
            left := i
            right := len(inorder) - i - 1
            
            if left > 0 {
                root.Left = buildTree(preorder[1:left+1], inorder[:left])
            }
            
            if right > 0 {
                root.Right = buildTree(preorder[1+left:], inorder[left+1:])
            }
            break
        }
    }
    
    return root
}
