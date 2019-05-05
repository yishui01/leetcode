/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    //两种方法，一种是一个一个比对，硬排，还有一种是将全部node放到数组中，sort后再组成链表输出,内存耗费较多
    //好吧我选择硬排,第一个和第二个比，比完后的结果再和第三个比，以此类推
    //一个个硬排的执行时间是108ms，参考8ms的答案进行优化，果然自己的答案也变成8ms了，要点在于要两两比对，而不是逐个比对
    
    lens := len(lists) 
    if lens == 0 {
        return nil
    }
    if lens == 1{
        return lists[0]
    }
    
    return merge(lists)

}

func merge(lists []*ListNode) *ListNode{
    lens := len(lists)
    half := lens/2
    if lens == 1 {
        return lists[0]
    } else if lens == 2 {
        return sortList(lists[0],lists[1])
    } else {
        return sortList(merge(lists[:half]),merge(lists[half:]))
    }
}

func sortList(a *ListNode, b *ListNode)*ListNode{
    if a == nil {
        return b
    } else if b == nil {
        return a
    }
    
    var head *ListNode
    if a.Val <= b.Val {
        head = a
        a = a.Next
    } else {
        head = b
        b = b.Next
    }
    tmp := &(head.Next)
    for a != nil && b != nil {
        if a.Val <= b.Val {
            *tmp = a
            a = a.Next
        } else {
            *tmp = b
            b = b.Next
        }
        tmp = &((*tmp).Next)
    }
    
    if a == nil {
        *tmp = b
    } else {
        *tmp = a
    }
    
    return head
    
}
/**
下面这个是未经优化版本的108ms的

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    //两种方法，一种是一个一个比对，硬排，还有一种是将全部node放到数组中，sort后再组成链表输出,内存耗费较多
    //好吧我选择硬排,第一个和第二个比，比完后的结果再和第三个比，以此类推
    
    lens := len(lists) 
    if lens == 0 {
        return nil
    }
    if lens == 1{
        return lists[0]
    }
    
    res := lists[0]
    
    for i:=1;i < lens; i++{
        res = sortList(res, lists[i])
    }
    
    return res
    
}

func sortList(a *ListNode, b *ListNode)*ListNode{
    if a == nil {
        return b
    } else if b == nil {
        return a
    }
    
    var head *ListNode
    if a.Val <= b.Val {
        head = a
        a = a.Next
    } else {
        head = b
        b = b.Next
    }
    tmp := &(head.Next)
    for a != nil && b != nil {
        if a.Val <= b.Val {
            *tmp = a
            a = a.Next
        } else {
            *tmp = b
            b = b.Next
        }
        tmp = &((*tmp).Next)
    }
    
    if a == nil {
        *tmp = b
    } else {
        *tmp = a
    }
    
    return head
    
}

**/
