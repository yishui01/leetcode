/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    //递归回溯
    if root == nil {
        return 0
    }
    res := 0
    helper(root, []int{}, &res)
    return res
    
}

func helper(node *TreeNode, tmpRes []int, res *int) {
    if node == nil {
        return 
    }
   
    if node.Left == nil && node.Right == nil {
        lens := len(tmpRes)
        dist := make([]int, lens)
        copy(dist, tmpRes)
        dist = append(dist, node.Val)
        newLens := lens + 1
        tmp := 0
        for k,v := range dist {
            tmp  += Pow(10,newLens-k-1) * v
        }
        *res = *res + tmp
        return 
    }
    
    if node.Left != nil {
        tmpRes = append(tmpRes, node.Val)
        helper(node.Left, tmpRes, res)
        tmpRes = tmpRes[:len(tmpRes)-1]
    }
    
    if node.Right != nil {
        tmpRes = append(tmpRes, node.Val)
        helper(node.Right, tmpRes, res)
        tmpRes = tmpRes[:len(tmpRes)-1]
    }
    
}

func Pow(num, index int) int {
    if num == 0 {
        return 0
    }
    if index == 0 {
        return 1
    }
    for i:=0; i < index-1; i++ {
        num *= 10
    }
    
    return num
}
