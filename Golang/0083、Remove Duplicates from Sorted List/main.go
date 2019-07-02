/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    tmpVal := head.Val
    res := head
    tmp := head
    
    head = head.Next
    for head != nil {
        if head.Val == tmpVal {
            tmp.Next = nil
        } else {
            tmp.Next = head
            tmp = tmp.Next
            tmpVal = tmp.Val
        }
        
        head = head.Next
    }
    
    return res
}
