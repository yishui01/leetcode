/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    //头插法
    dummy := &ListNode{}
    
    for head != nil {
        next := head.Next
        head.Next = dummy.Next
        dummy.Next = head
        head = next
    }
    
    return dummy.Next
}
