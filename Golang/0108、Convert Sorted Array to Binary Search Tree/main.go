/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    //被题目给误导了，我以为只要把nums对折就行了，后来答案错的一塌糊涂，才想起是要平衡树，o(╯□╰)o
    //分治找中点作为根结点
   
    res := recursion(nums)
    return res
    
    //下面为解法二
    //lens := len(nums)
    //res := recursion(0, lens, nums)
    // return res
}

//下面的是最原始的传的nums
func recursion(nums []int) *TreeNode {
    lens :=len(nums)
    if lens == 0 {
        return nil
    }
    mid := lens/2
    node := new(TreeNode)
    node.Val = nums[mid]
    if mid > 0 {
        node.Left = recursion(nums[:mid])
    }
    if mid < len(nums)-1 {
        node.Right = recursion(nums[mid+1:])
    }
    return node
}

//解法二
// func recursion(start, end int, nums []int) *TreeNode {
//     if end-start <= 0 {
//         return nil
//     }
//     mid := (end+start)/2
//     node := new(TreeNode)
//     node.Val = nums[mid]
//     if mid > 0 {
//         node.Left = recursion(start, mid, nums)
//     }
//     if mid < end-1 {
//         node.Right = recursion(mid+1, end, nums)
//     }
//     return node
// }



