/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func sortedListToBST(head *ListNode) *TreeNode {
  //快慢指针找中点，分治,//无副作用方案，不会更改链表
    if head == nil {
        return nil
    }
    return helper(head, nil)
}

func helper(head, res *ListNode) *TreeNode {
    midList := findMid(head, res) //第二个参数要传res，不能直接写nil，不然分治时左右子树的的终点不对
    node := new(TreeNode)
    if midList != nil {
        node.Val = midList.Val
    } else {
        return nil
    }
    
    //分治
    node.Left = helper(head, midList)
    node.Right = helper(midList.Next, res) //这里很容易错，就是第二个参数不能直接写成nil，那样就会导致左子树在生成他的右子树时，终点不对，实际上终点应该为之前的mid, 而不是根节点右子树的终点nil
    
    return node
    
}

//head为头指针，res为链表终点的Next
func findMid(head, res *ListNode) *ListNode {
    if head == nil || head == res {
        return nil
    }
    if head.Next == res {
        return head
    }
    
    fast := head
    slow := head
    
    for fast != nil &&fast != res && fast.Next != nil && fast.Next != res {
        fast, slow  = fast.Next.Next, slow.Next
    }
   
    return slow
}







// func sortedListToBST(head *ListNode) *TreeNode {
//     //解题思路：双指针找链表的中点，分治
    
//    return helper(head, nil)
    
// }

// func helper(head *ListNode, res *ListNode) *TreeNode {
//     if head == nil || head == res {
//         return nil
//     }
//      //开始找中点
//     mid := findMid(head, res)
//     node := new(TreeNode)
//     if mid == nil { 
//        return nil  //中点为空直接返回
//     } else {
//         node.Val = mid.Val //否则作为根节点
//     }
    
//     //分治左右子树,除去中点
//     node.Left = helper(head, mid)
//     node.Right = helper(mid.Next, res) //这里之前第二个参数传的是nil，怎么输入都不对，后台模拟了一遍才发现不能传nil，因为左子树的终点并不是nil，左子树的终点应该是之前的mid
    
//     return node
// }
 
// //找出从head到res的中点，当快指针或者慢指针的next为res时，中点就是slow
// func findMid(head *ListNode, res *ListNode) *ListNode {
//     if head == res { //head就是终点，终点是不能到达的，只有当quick或者slow的Next指向res的状态才是合法的，所以这里直接返回nil
//         return nil
//     }
//     if head.Next == res {  //head的下一个就是终点，那目前已经是达成成就的状态，直接返回
//         return head
//     }
//     quick := head
//     slow := head
//     for quick.Next != res {
//         if slow.Next != res {
//             slow = slow.Next
//         } else {
//             return slow
//         }
        
//         if quick.Next != res {
//             quick = quick.Next
//         } else {
//             return slow
//         }
        
//         if quick.Next != res {
//             quick = quick.Next
//         } else {
//             return slow
//         }
//     }
    
//     return slow
// }




