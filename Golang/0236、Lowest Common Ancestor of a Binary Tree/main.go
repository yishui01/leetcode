/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */



/*
//解法一：左右子树分别递归，看p和q是否在同一个子树上，在的话就返回这个子树的根节点，否则返回当前的根节点（因为p和q分散在当前树的左右两边）
p和q有两种情况 
一、在相同子树上
二、在不同子树上

从根节点遍历，递归向左右子树查询节点信息
递归终止条件：如果当前节点为空或等于p或q，则返回当前节点

递归遍历左右子树，如果左右子树查到节点都不为空，则表明p和q分别在左右子树中，因此，当前节点即为最近公共祖先；
如果左右子树其中一个不为空，则返回非空节点。
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
 if root == nil || root == p || root ==q {
     return root
 }
 left := lowestCommonAncestor(root.Left,p,q)
 right := lowestCommonAncestor(root.Right,p,q)
 if left != nil && right != nil {
     return root
 }
 if left == nil {
     return right
 }
 return left
}



//解法二 回溯
//  func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//      var res *TreeNode
//      helper(root,p,q,&res)
//      return res
// }

// func helper(root,p,q *TreeNode,res **TreeNode) bool {
//     if root == nil {
//         return false
//     }
    
//     left,right,mid := 0,0,0
//     if helper(root.Left,p,q,res) {
//         left++
//     }
    
//     //看是否已经找到结果
//     if *res != nil {
//         return false
//     }
    
//     if helper(root.Right,p,q,res) {
//         right++
//     }
    
//     if root == p || root == q {
//         mid++
//     }
    
//     if mid + left + right >= 2{
//         *res = root
//     }
    
//     return mid + left + right > 0
// }
