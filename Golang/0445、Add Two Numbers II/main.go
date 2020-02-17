/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//先把两个链表翻转，再相加    (还有一种不翻转链表的情况就是建立两个栈，根据两个栈计算数值并建表)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var r1,r2 *ListNode
    r1,r2 = reverse(l1),reverse(l2)
    isAdd := 0 //是否进位

    dummy := &ListNode{}
    now := dummy

    for r1 != nil || r2 != nil || isAdd != 0 {
        r1V,r2V := 0,0
        if r1 != nil {
            r1V = r1.Val
        }
        if r2 != nil {
            r2V = r2.Val
        }

        sum := r1V + r2V + isAdd
        if sum >= 10 {
            isAdd = 1
            sum -= 10
        } else {
            isAdd = 0
        }
        now.Next = &ListNode{Val:sum,Next:nil}
        now = now.Next

        if r1 != nil {
            r1 = r1.Next
        }
        if r2 != nil {
            r2 = r2.Next
        }
    }

    res := reverse(dummy.Next)

    return res

}

func reverse(head *ListNode)*ListNode {
    var new *ListNode    
    for head != nil {
       preNext := head.Next
       head.Next = new
       new = head
       head = preNext
    }

    return new
}
