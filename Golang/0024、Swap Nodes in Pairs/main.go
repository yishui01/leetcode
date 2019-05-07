/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    //硬换就行，两个一换，换完之后将上一个的尾部指向这两个的头部，再将这两个的尾部指向下一个node，并将尾部作为下一次的上一个尾部
    var res,prev *ListNode
    isStart := 0
    for head != nil  && head.Next != nil {
        isStart = 1
        node1 := head
        node2 := head.Next
        //第一步、将前一个节点指向后一个节点的Next，也就是第三个节点的地址
        node1.Next = node2.Next
        node2.Next = node1
        //第二步、此时node2在前，挂到前一个节点的尾部
        if res == nil {
            res = node2
            prev = node1
        } else {
            prev.Next = node2
            prev = node1
        }
        head = node1.Next
    }
    
    if isStart == 0 {
        return head
    }
    
    if head != nil {
        prev.Next = head
    }
    
    return res
    
}