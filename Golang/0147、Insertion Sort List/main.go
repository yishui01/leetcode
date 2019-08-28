/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    res := head
    head = head.Next
    tmp := res
    tmp.Next = nil  //记住，每个添加到res链表中的尾部节点都要清空Next，否则可能成环
    var prev *ListNode
    
    for head != nil {
        nextHead := head.Next
        isFind := false
        for tmp != nil {
            if tmp.Val > head.Val {
                isFind = true
                //找到了
                if prev == nil {
                    //比头节点都小,直接添加到头部
                    prevRes := res
                    res = head
                    head.Next = prevRes
                } else {
                    //中间的某个节点
                    prev.Next = head
                    head.Next = tmp
                    prev = nil
                }
                break
            } else {
                prev = tmp
                tmp = tmp.Next
            }
        }
        
        if !isFind {
            //没找到的话,说明比所有节点都大，挂到尾部
            prev.Next = head
            head.Next = nil  //清空尾部
        }
        prev = nil
        tmp = res
        head = nextHead
    }
    
    return res
}

//为什么大佬的解法总是这么清晰简洁？

// func insertionSortList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	dummy := &ListNode{-1, head}
// 	prev := head
// 	cur := head.Next
// 	for cur != nil {
// 		if cur.Val >= prev.Val {
// 			prev = cur
// 			cur = cur.Next
// 		} else {
// 			prev.Next = cur.Next
// 			p1 := dummy.Next
// 			p2 := dummy
// 			for p1 != nil {
// 				if p1.Val < cur.Val {
// 					p2 = p1
// 					p1 = p1.Next
// 				} else {
// 					break
// 				}
// 			}
// 			cur.Next = p1
// 			p2.Next = cur
// 			cur = prev.Next
// 		}
// 	}
// 	return dummy.Next
// }
