/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    //封装一个函数，不停的换就行，如果链表长度为1的话直接返回
    if head == nil || head.Next == nil {
        return head
    }
    
    var end *ListNode
    var lens int
    tmp := head
    for tmp != nil {
        end = tmp
        lens++
        tmp = tmp.Next
    }
    
    //超时后发现这个交换过程是不是可以省略一些重复的轮回啊，取模了解一下,例如长度为2的链表，交换0次和交换2次、4次是一样的结果
    k%=lens
    for i:=0; i < k; i++ {
        head,end = rotate(head, end)
    }
    
    return head
}

func rotate(head, end *ListNode) (*ListNode, *ListNode){
    var preEnd,e *ListNode
    tmp := head
    for tmp != nil {
        preEnd = e
        e = tmp
        tmp = tmp.Next
    }
    
    //end的前一个指向nil
    preEnd.Next = nil
    //原来的end指向head
    end.Next = head
    return end,preEnd
    
}
