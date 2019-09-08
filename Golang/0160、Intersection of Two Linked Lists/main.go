/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
  //解法二：分别同时遍历两个链表，A遍历完了之后接到B的head，B遍历完了之后接到A的head，最终他们会在交叉点相遇，因为他们即便第一次短
//链表先经过了交叉点，导致没能相遇在交叉点，那么短链表遍历完之后转到长链表开始遍历时他还是要还回来的，反之长链表遍历完到短链表就会上来，差值追平
//的时候一定就是相交的点的时候，当然也有可能两者没有相交，最后 相等的点就是 nil
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    
    if headA == nil || headB == nil {
        return nil
    }
    
    l1 := headA
    l2 := headB
    for l1 != l2 {
        if l1 != nil{
           l1 = l1.Next
        } else {
           l1 = headB
        }
        if l2 != nil {
           l2 = l2.Next
        } else {
           l2 = headA
        }
    }
    
    return l1
   
}


//   //解法一：假设两个链表长度相等，那么逐个比对即可，可是存在不相等的情况，那就把长的变得和短的一样长，后移后开始比对
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
//     //跪了跪了，现在连easy的题都做不出来了么，越做越回去了啊 ╮(╯▽╰)╭
//     l1 := getLength(headA)
//     l2 := getLength(headB)
    
//     if l1 == 0 {
//         return headA
//     }
//     if l2 == 0 {
//         return headB
//     }
    
//     if l1 != l2 {
//         diff := 0
//         if l1 > l2 {
//             //后移headA
//             diff = l1 - l2
//             for diff != 0 { //后移过程中无需比对，因为长出的部分怎么也不可能是相交的点
//                 headA = headA.Next
//                 diff--
//             }
//         } else {
//             //后移headB
//              diff = l2-l1
//             for diff != 0 {
//                 headB = headB.Next
//                 diff--
//             }
//         }
//     }
    
//     for headA != headB {
//         headA = headA.Next
//         headB = headB.Next
//     }
//     return headA
// }

// func getLength(head *ListNode) int {
//     num := 0
//     for head != nil {
//         num++
//         head = head.Next
//     }
//     return num
// }
