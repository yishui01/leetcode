/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    //这题挂了，不会做。
    //迭代思路：迭代root的左子树，一直迭代下去，每次迭代如果该节点有左孩子，找到该左孩子的最右节点，将最右节点的右孩子设置为迭代节点的右孩子，将迭代节点左孩子置空，右孩子设置为左孩子，
    //由于是一直将左边置空，挂到右边，所以迭代方式是一直往右，root = root.Right
    for root != nil {
        if root.Left != nil {
            rootLeft := root.Left 
            for rootLeft.Right != nil{
                rootLeft = rootLeft.Right
            }
            rootLeft.Right = root.Right
            root.Right = root.Left
            root.Left = nil
        }
        root = root.Right
    }
}

/*
//看下大佬的递归版本，递归实际上是自底向上的，保留右边为一个变量，替换右边为左边，将左边置空，迭代到最右，将最右一个节点的右孩子设置为之前保存的右边的变量，自底向上
class Solution {
public:
    void flatten(TreeNode *root) {
        if (!root) return;
        if (root->left) flatten(root->left);
        if (root->right) flatten(root->right);
        TreeNode *tmp = root->right;
        root->right = root->left;
        root->left = NULL;
        while (root->right) root = root->right;
        root->right = tmp;
    }
};
*/
