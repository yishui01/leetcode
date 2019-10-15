/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteNode(node *ListNode) {
    if node == nil {
        return 
    }
   node.Val = node.Next.Val
   node.Next = node.Next.Next
    
}


// func deleteNode(node *ListNode) {
//     if node == nil {
//         return 
//     }
//     var pre *ListNode
    
//     for node.Next != nil {
//         pre = node
//         node.Val = node.Next.Val
//         node = node.Next
//     }
//     if pre != nil {
//        pre.Next = nil
//     }
    
// }

