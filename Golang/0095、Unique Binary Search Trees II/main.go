/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    //这题主要是要去重复，直接递归回溯还不行。
    //因此要考虑一个巧妙的办法，用二叉树的性质，以当前i为root，小于i的作为其左子树，大于i为作为其右子树,实际上这个递归不是回溯了，这是分治,先拆分，再聚合
    if n == 0 {
        return nil
    }
    return recursion(1,n)
}


func recursion(start, end int)[]*TreeNode{
    res := make([]*TreeNode, 0)
    //先确定边界条件
     //此时没有数字，将 null 加入结果中
    if start > end {
        return append(res, nil)
    }
    
   //只有一个数字，当前数字作为一棵树加入结果中
    if start == end{
        tree := new(TreeNode);
        tree.Val=start
        res = append(res, tree)
        return res;
    }
    
     //尝试每个数字作为根节点
    for i:=start; i <= end; i++ {
         //得到所有可能的左子树
        leftTrees := recursion(start, i - 1);
         //得到所有可能的右子树
        rightTrees := recursion(i + 1, end);
        //左子树右子树两两组合
        for _,left:=range leftTrees {
            for _,right:=range rightTrees {
                root := new(TreeNode)
                root.Left = left
                root.Right = right
                root.Val = i
                res = append(res, root)
            }
        }

    }
    
    return res
    
}
