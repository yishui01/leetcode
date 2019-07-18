/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recoverTree(root *TreeNode)  {
    //考虑BST的性质，左子树的最右结点 < 当前结点 < 右子树的最小结点
    //只要不满足当前性质，就直接交换两个结点，每次只进行一次交换操作，交换完直接return ，外层通过for循环进行下一轮检查，直到左右子树全部完成
    for !isBst(root) {
        
    }
    
}

func isBst(node *TreeNode) bool {
    if node == nil {
        return true
    }
    //先检查左子树
    if node.Left != nil {
        next := node.Left
        for next.Right != nil {
            next = next.Right
        }
        if next.Val > node.Val {
            tmp := node.Val
            node.Val = next.Val
            next.Val = tmp
            return false
        }
        if !isBst(node.Left) {
            return false
        }
    }
    
    if node.Right != nil {
        next := node.Right
        for next.Left != nil {
            next = next.Left
        }
        if next.Val < node.Val {
            tmp := node.Val
            node.Val = next.Val
            next.Val = tmp
            return false
        }
        if !isBst(node.Right) {
            return false
        }
    }
    
    return true
    
}



//本解法不符合O(n)的空间复复杂度
// func recoverTree(root *TreeNode)  {
//     //那就先来一个O(n)的，原来二叉树还可以这样给结点赋值，就是把一个二叉树压平，
//     //压平就是用中序遍历将二叉树的node存好，再将值也存好，再将值排好序，再逐个赋值给node
//     nodeArr := make([]*TreeNode, 0)
//     nums := make([]int, 0)
//     recursion(root, &nodeArr, &nums)
//     sort.Ints(nums)
//     for i:=0; i < len(nums); i++ {
//         nodeArr[i].Val = nums[i]
//     }
// }

// func recursion(node *TreeNode, nodeArr *[]*TreeNode, nums *[]int) {
//     if node == nil {
//         return 
//     }
//     recursion(node.Left, nodeArr, nums)
//     *nodeArr = append(*nodeArr,node)
//     *nums = append(*nums, node.Val)
//     recursion(node.Right, nodeArr, nums)
// }
