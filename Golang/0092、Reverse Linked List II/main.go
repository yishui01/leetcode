/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    //单向链表翻转,三指针法
    //记录翻转结点的前一个结点
    //翻转指定区域
    //到了最后一个结点的时候，将之前记录的翻转结点的前一个结点的next指向这个结点，再将翻转结点的第一个结点的Next指向当前结点的Next
    //如果存在 翻转结点的前一个结点，那么就返回头结点，否则返回翻转区域的最后一个结点
    
    now := 1
    
    var res *ListNode
    var prev *ListNode
    var next *ListNode
    var current *ListNode
    
    var beginPre *ListNode
    var begin *ListNode
    
    if head.Next == nil || m == n {
        return head
    }
    
    if m!=1 {
        res = head
    }
    
    current = head
    
    for current != nil {
        if now > n {
            break
        }
        
        next = current.Next
        
        if now == m-1 {
            beginPre = current 
        }else if now >= m && now <= n {
            if now == m {
                 begin = current
            }
            
            if prev != nil {
                current.Next = prev
                prev = current
                current = next
            } else {
                prev = current
            }
            
            if now == n {
                if beginPre != nil {
                    beginPre.Next = prev
                    begin.Next = next
                } else {
                    begin.Next = next
                    res = prev
                }
                break
            }
        }
        now++
        current = next
        
    }
    
    return res
    
}
