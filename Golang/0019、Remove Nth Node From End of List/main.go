/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    //先遍历一遍链表，把每个节点的头指针都保留下来，最后去除target，拼接后返回
    if n == 0 {
        return head
    }
    maps:= make(map[int]*ListNode)
    tmp := head
    i:=1
    for tmp != nil {
        maps[i] = tmp
        i++
        tmp  = tmp.Next
    }
    //i-1节点数，去掉倒数第n个节点，索引为i-1-(n-1)
    delIndex := i-1-(n-1)
    p,ok := maps[delIndex]
    if !ok {
        return head
    }
    if p == head {
        return head.Next
    }
    p2,_ := maps[delIndex-1]
    p3,_ := maps[delIndex+1]
    p2.Next = p3
    return head
    
    
}
