/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

//解法二：利用二叉树的大小性质dfs
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     if root == nil {
         return nil
     }
     if root.Val > p.Val && root.Val > q.Val{
         //如果root比两个节点的值都大，说明两个节点都在root的Left子树上
         return lowestCommonAncestor(root.Left,p,q)
     }
     
     if root.Val < p.Val && root.Val < q.Val{
         //如果root比两个节点的值都小，说明两个节点都在root的Right子树上
         return lowestCommonAncestor(root.Right,p,q)
     }
     
     //否则就是说p和q一个在左一个在右，当前root就是最近公共祖先
     return root
}

//解法三：解法一的优化版本，p和q可能在同一子树中也可能不在同一子树中
//  func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//      if root == nil || root == p || root ==q {
//          return root
//      }
//      left := lowestCommonAncestor(root.Left,p,q)
//      right := lowestCommonAncestor(root.Right,p,q)
//      if left != nil && right != nil {
//          return root
//      }
//      if left == nil {
//          return right
//      }
//      return left
// }



// 下面这解法还是不行，重复操作太多了
//解法一：dfs分别找p和q是否在当前子树中  
//思路：在当前的root节点为根的树中，先找p再找q，如果都找到了，说明root是合格的，那么继续看root的Left是不是合格的，如果是的就继续递归Left节点，
// 如果Left不合格，就递归Right节点，返回最终合法的那个
//  func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//      if root == nil {
//          return nil
//      }
//      if helper(root,p) && helper(root,q) {
//          resLeft := lowestCommonAncestor(root.Left,p,q)
//          if resLeft != nil {
//              return resLeft
//          }
//          resRight := lowestCommonAncestor(root.Right,p,q)
//          if resRight != nil {
//              return resRight
//          }
//          return root
//      }
//      return nil
// }

// //找p点是否在root为根节点树中
// func helper(root,p *TreeNode) bool {
//     if root == nil {
//         return false
//     }
    
//     if root.Val == p.Val {
//        return true
//     }
    
//     return helper(root.Left,p) || helper(root.Right,p)
// }
