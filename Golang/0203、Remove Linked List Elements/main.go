/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{Next:head}
    p1 := dummy
    p2 := head
    
    for p2 != nil {
        if p2.Val == val {
              p1.Next = p2.Next
              p2 = p2.Next
        } else {
            p1 = p2
            p2 = p2.Next
        }
    }
    
    return dummy.Next
}


/**

大佬的解法真TM 丧心病狂
func removeElements(head *ListNode, val int) *ListNode {
    if head==nil{
        return nil
    }
    
    if head.Val==val{
        return removeElements(head.Next,val)
    }else{
        head.Next=removeElements(head.Next,val)
        return head
    }
}
***/
