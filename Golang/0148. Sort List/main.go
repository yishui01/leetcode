/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    //归并排序
    //归并的思想就是分治
    //用快慢指针将链表拆成两段，再用递归归并
    if head == nil || head.Next == nil {
        return head
    }
    fast,slow,prev := head, head, head
    for fast!= nil && fast.Next != nil {
        prev = slow
        fast = fast.Next.Next
        slow = slow.Next
    }
    prev.Next = nil
    return merge(sortList(head),sortList(slow))
}

//将两个链表按从小到大排好序后合成一个返回
func merge(l1, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
         return l1
    }
    if l1.Val <= l2.Val {
        l1.Next = merge(l1.Next, l2)
        return l1
    } else {
        l2.Next = merge(l2.Next, l1)
        return l2
    }
   
}
