/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    //O（n）时间+O（1）空间，找到中点，将前半段逆序
    //快慢指针找到中点
    
    if head == nil || head.Next == nil {
        return true
    }
    
    quick,slow := head,head
    var prev,mid *ListNode
    count := 0
     
    for quick != nil && quick.Next != nil {
        prev = slow
        quick = quick.Next.Next
        slow = slow.Next
        count++
    }
    
    if quick == nil {
         //偶数时 prev指向中间两个点的前一个点,slow指向后一个点
        //这个时候就要比对下中间两个点是否相等，不然就漏了
        if slow.Val != prev.Val {
            return false
        }
        mid = slow
        count--
        mid = mid.Next
    } else {
         //奇数时 slow指向中点
        mid = slow 
        mid = mid.Next
    }
    
    if count == 0 { //只有两个节点的情况
        return head.Val == head.Next.Val
    }
    
    //逆序前半段
    nowNode := head
    var prevNode *ListNode
    
    for count !=0 {
        tmpNext := nowNode.Next
        nowNode.Next = prevNode
        
        prevNode = nowNode
        nowNode = tmpNext
        count--
    }
    
    //再开始进行比对
    for prevNode!=nil || mid != nil {
        if prevNode.Val != mid.Val {
            return false
        }
        prevNode = prevNode.Next
        mid = mid.Next
    }
    
    return true
    
}
