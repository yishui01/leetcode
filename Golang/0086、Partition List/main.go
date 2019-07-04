/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    //把大于等于的拆出来组成一个新链表
    if head == nil {
        return nil
    }
    var minList *ListNode //把小的拆到这里来
    var minTmp *ListNode
    
    var newList *ListNode //把大的挂到这里
    var newTmp *ListNode
    
    var res *ListNode
    
    for head != nil {
        if head.Val >= x {
            if newList == nil {
                newList = head
                newTmp = head
            } else {
                newTmp.Next = head
                newTmp = newTmp.Next
            }
        } else {
            if minList == nil {
                minList = head
                minTmp = head
            } else {
                minTmp.Next = head
                minTmp = minTmp.Next
            }
        }
        head = head.Next
    }
    
    if newTmp != nil {
        newTmp.Next = nil
    }
    
    if minList != nil {
        res = minList
        minTmp.Next = newList
    } else {
        res = newList
    }
    
    return res
    
}
