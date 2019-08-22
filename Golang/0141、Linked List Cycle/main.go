/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    //Easy的题，但是我就是做不出来，快慢指针，一个走一步一个走两步，如果有环最终会相遇，蓝瘦，香菇，为什么我想不出来
    quick := head
    slow := head
    
    for quick != nil && quick.Next != nil {
        quick = quick.Next.Next
        slow = slow.Next
        if quick != nil && quick.Val == slow.Val {
            return true
        }
    }
    
    return false
}
