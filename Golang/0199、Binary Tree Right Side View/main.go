/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    //右视图,思路：层级遍历，先遍历右子树，将新的节点添加到res中，接下来判断左子树的时候有个 level > len(*res)
    //如果右子树已经填充好了，那左子树的将被忽略，如果否则将填充左子树
    res := make([]int,0)
    dfs(root,&res,1)
    return res
}

func dfs(node *TreeNode,res *[]int,level int){
    if node == nil {
        return 
    }
    if level > len(*res){
        *res = append(*res,node.Val)
    }
    dfs(node.Right, res, level+1)
    dfs(node.Left, res, level+1)
}
