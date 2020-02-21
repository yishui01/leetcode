/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
   find := false
   return helper(root,nil,key,0,&find)
}


func helper(root,parent *TreeNode,key,dic int,find *bool) *TreeNode {
     if root == nil || *find {
        return root
    }
    newRoot := root 
    if root.Val == key {
        *find = true
        switch{
            case root.Left == nil && root.Right == nil:   //叶子节点
                newRoot = nil
            case root.Left == nil && root.Right != nil:  //只有右节点
                newRoot = root.Right
            case root.Left != nil && root.Right == nil:  //只有左节点
                newRoot = root.Left
            default: //两个都有,直接找后继，就是right的LeftMost
                r := root.Right
                p := root
                step := 0
                for r.Left != nil {
                    p = r
                    r = r.Left
                    step++
                }
                //先把后继的子节点处理了,后继是右边的leftMost，所以没有左子节点，只要处理右子节点
                if step == 0 { //代表这个parent的右节点没有left了，要取代parent的右边
                    p.Right = r.Right
                } else {  //有left，指向parent的左边
                    p.Left = r.Right
                }

                //再把r换到删除的那个节点的位置
                r.Left,r.Right = root.Left,root.Right
                newRoot = r
        }
        //看下parent是不是nil,是的话不用管，直接返回新节点newRoot
        if parent != nil {
            //处理parent的指向
            if dic == 0 {
                //左边
                parent.Left = newRoot
            } else {
                parent.Right = newRoot
            }
        }
    }

    if !*find {
        if key < newRoot.Val {
            newRoot.Left = helper(newRoot.Left,newRoot,key,0,find)
        } else {
            newRoot.Right = helper(newRoot.Right,newRoot,key,1,find)
        }
    }

    return newRoot

}

