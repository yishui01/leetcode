/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    //一段一段的转就行,将node存到数组中，存到k个了就反转数组，拼成个新链表挂到res上
    if head == nil {
        return nil
    }
    if k == 1 {
        return head
    }
    var arr [1300]*ListNode
    var res,prev *ListNode
   
    i := 0
    for head != nil {
        if i == k{
             i = 0
             reverse(k,&arr, &res, &prev)
        }
        arr[i] = head
        head = head.Next
        i++
    }
    if i == k {
        //这是刚好剩k个节点，然后head为nil了，没挂上去，所以这里挂上去
        reverse(k,&arr, &res, &prev)
    } else {
        if prev == nil {
            //这里是一次反转都没有，总共都不够k个
            return arr[0]
        }
        //剩余的不够k个，放数组里了，直接挂到res上就行，不要反转了
        prev.Next = arr[0]
    }
    
    return res
}

//反转数组，把数组反转后拼成个链表挂到res中
func reverse(k int, arr *[1300]*ListNode, res **ListNode, prev **ListNode){
    for j:=0; j < k/2; j++{
            tmp := arr[j]
            arr[j] = arr[k-1-j]
            arr[k-1-j] = tmp
        }
        
        for z:=0;z<k-1;z++ {
            arr[z].Next = arr[z+1]
        }
        arr[k-1].Next = nil
        if *prev == nil {
            *res = arr[0]
            *prev = arr[k-1]
        } else {
            (*prev).Next = arr[0]
            *prev = arr[k-1]
        }
}








