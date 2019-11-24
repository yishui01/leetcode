/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return head
    }
    
    //用一个指针保存分界点，分界点为奇数节点的最后一个节点，他的Next就是偶数节点的第一个节点
    tmpHead := head.Next.Next
    split := head
    prev := head.Next
    i := 1
    for tmpHead != nil {
        if i & 1 == 0 {
            //偶数节点，直接迭代指针即可
            tmpHead = tmpHead.Next
            prev = prev.Next
        } else{
            //奇数节点
            
            //向奇数段尾部添加当前节点
            splitNext := split.Next //保存原来的Next
            split.Next = tmpHead //将当前节点挂上去
            split = split.Next //后移指针
            
            //整理刚刚挂上去的节点的Next
            tmpNext := tmpHead.Next 
            tmpHead.Next = splitNext
            
            //将挂上去的节点和之前和他相连的前一个节点拆分
            prev.Next = tmpNext
            
            //迭代指针
            tmpHead = tmpNext
        }
        
        i++
    }
    
    return head
}
