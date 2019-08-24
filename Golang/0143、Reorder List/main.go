/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
/**
查看答案看到大佬的解法就是6
1、获取中点指针
2、将原链表切分成两部分，记得打断中点的连接
3、反转node2
4、插花合并两个链表，注意前半段数量比后半段数量少一个的情况
**/
    fast := head
    slow := head
    var pre *ListNode
    
    for fast != nil  && fast.Next != nil {
        pre = slow
        fast,slow = fast.Next.Next, slow.Next
    }
    if pre != nil { //注意要打断前半段和后半段的连接
        pre.Next = nil
    }
    if slow == nil || slow == head { //==head是代表只有一个节点的情况，直接return
        return 
    }
    //slow指向的是中点，不用考虑奇偶性，从slow到末端直接全部反转就行，因为效果是一样的 
    cur := slow.Next
    slow.Next = nil
    for cur != nil {
        tmp := cur.Next
        cur.Next = slow
        slow = cur
        cur = tmp
    }
    //反转完成之后slow为最后一个节点，那就是头结点
    
    //再将反转后的后半段链表插花到前半段中
    for head != nil {
        headTmp := head.Next
        head.Next = slow
        if headTmp != nil { //前半段数量比后半段数量少一个的情况
            slowTmp:=slow.Next
            slow.Next = headTmp
            slow = slowTmp
        }
        head = headTmp
    }
    
    return 
    
    
}
