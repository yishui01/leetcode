/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    res := make([]string,0)
    helper(root,"",&res)
    return res
}

func helper(node *TreeNode,str string,res *[]string){
    if node == nil{
        return 
    } 
    if str == ""{
        str = strconv.Itoa(node.Val)
    } else {
        str += "->"+strconv.Itoa(node.Val)
    }
    if node.Left == nil && node.Right == nil {
         *res = append(*res,str)
         return
    }
    helper(node.Left,str,res)
    helper(node.Right,str,res)
}
