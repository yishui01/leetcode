/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//中序遍历
type BSTIterator struct {
    //直接用栈,先将左子节点全部存入
    Stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    obj := BSTIterator{}
    for root != nil {
        obj.Stack = append(obj.Stack,root)
        root = root.Left
    }
    return obj
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    //思路就是将左子节点全部存入，存入的节点弹出并起作用之后就把指向他的Right，再把他的Right的左子树遍历到的节点全部存入Stack
    p := this.Stack[len(this.Stack)-1]
    this.Stack = this.Stack[:len(this.Stack)-1]
    res := p.Val
    if p.Right != nil {
        p = p.Right
        for p != nil {
            this.Stack = append(this.Stack,p)
            p = p.Left
        }
    }
    return res
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return len(this.Stack) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
