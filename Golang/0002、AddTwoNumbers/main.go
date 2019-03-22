/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//先把两个链表中的值保存到两个slice中，再执行拼接
	var a, b []int

	for tmp := l1; tmp != nil; tmp = tmp.Next {
		a = append(a, tmp.Val)
	}

	for tmp := l2; tmp != nil; tmp = tmp.Next {
		b = append(b, tmp.Val)
	}
	var length int
	if len(a) >= len(b) {
		length = len(a)
	} else {
		length = len(b)
	}
    
	//接下来开始拼接新slice
	var c []int
	var is_overflow int = 0
	var tmpa, tmb int = 0, 0
    
    for i := 0; i < length;i++{
		if len(a) <= i {
			tmpa = 0
		} else {
			tmpa = a[i]
		}

		if len(b) <= i {
			tmb = 0
		} else {
			tmb = b[i]
		}

        if tmpa + tmb+is_overflow >=10 {
            c = append(c, tmpa + tmb+is_overflow-10)
            is_overflow = 1
        } else {
            c = append(c, tmpa + tmb+is_overflow)
            is_overflow = 0
        }
	
	}
    if is_overflow > 0{
        fmt.Println("还有进位",is_overflow)
        c = append(c,is_overflow)
    }
    
	//c拼接完成之后，组合成链表
    var point, pre_point, head *ListNode
	for _, v := range c {
      point = &ListNode{
			Val:  v,
			Next: nil,
		}
        if pre_point !=nil {
            pre_point.Next=point
        }
		pre_point = point
        if head == nil {
            head = pre_point
        }
	}
    fmt.Println(c);
	return head

}
