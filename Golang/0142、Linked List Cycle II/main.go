/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    //看了大佬的答案，深深的体会到什么叫做智商是硬伤
    //解析：图就不话了，参考的这个解答https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/tu-jie-guo-cheng-by-yzzz0_0/
    //2(AB+BD) = AB + n(BCDEB) + BD
    //AB+BD=n(BCDEB)=nBD + nDEB
    //AB=(n-1)BCD+nDEB
    
    //至此得出 AB的长度为n-1圈环的长度+DEB的长度，此时的从DEB的起点开始走，最后相遇的点比为环的起点
    
    fast := head
    slow := head
    
    var res *ListNode
    
    for fast!= nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            //相遇了
            break
        }
    }
   if fast == nil || fast.Next == nil {
        return res
    }
    slow = head
    for fast != slow{
        fast = fast.Next
        slow = slow.Next
    }
    return fast
    
}
