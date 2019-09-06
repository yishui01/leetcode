//之前是想的用链表来排序存储节点，最小值就是链表的头部元素
//然而看了大佬的答案之后给跪了，其实可以用切片来存储最小值，如果新元素小于最小值就添加到切片的末尾，末尾就是最小值，如果大于就不管
//pop时如果pop的值和最小值相等就把min中的元素也pop了，不然就不用pop，跪了 20ms

type MinStack struct {
    Nums []int
    Min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    this.Nums = append(this.Nums, x)
    if len(this.Min) == 0 || this.Min[len(this.Min)-1] >= x {
        this.Min = append(this.Min, x)
    }
    
}


func (this *MinStack) Pop()  {
    if len(this.Nums) > 0  {
        val := this.Nums[len(this.Nums)-1]
        this.Nums = this.Nums[:len(this.Nums)-1]
        if this.Min[len(this.Min)-1] == val {
            this.Min = this.Min[:len(this.Min)-1]
        }
    }
}


func (this *MinStack) Top() int {
    return this.Nums[len(this.Nums)-1]
}


func (this *MinStack) GetMin() int {
    return this.Min[len(this.Min)-1]
}


/********************以下为本人的链表排序存储最小值的方案，很尬，60ms

//栈用一个数组就行，最小值用一个链表？
// type Node struct {
//     Val int
//     Next *Node
// }

// type MinStack struct {
//     Nums []int
//     Head *Node
// }

// /** initialize your data structure here. */
// func Constructor() MinStack {
//     return MinStack{
//         Nums:make([]int, 0),
//     }
// }


// func (this *MinStack) Push(x int)  {
//     this.Nums = append(this.Nums, x)
//     //添加到链表中
//     dummy := &Node{
//         Val:-1,
//     }
//     dummy.Next = this.Head
//     p1 := dummy
//     p2 := this.Head
//     for p2 != nil && p2.Val < x {
//         p1 = p2
//         p2 = p2.Next
//     }
//     p1.Next = &Node{
//         Val:x,
//         Next:p2,
//     }
//     this.Head = dummy.Next
    
// }


// func (this *MinStack) Pop()  {
//     if len(this.Nums) == 0 {
//         return 
//     }
//     val := this.Nums[len(this.Nums)-1]
//     this.Nums = this.Nums[:len(this.Nums)-1]
//     //从链表中删掉
//     dummy := &Node{
//         Val:-1,
//     }
//     dummy.Next = this.Head
//     p1 := dummy
//     p2 := this.Head
//     for p2 != nil && p2.Val != val {
//         p1 = p2
//         p2 = p2.Next
//     }
//     if p2 != nil {
//         p1.Next = p2.Next
//     } else {
//         panic("not val in slice")
//     }
//     this.Head = dummy.Next
// }


// func (this *MinStack) Top() int {
//     if len(this.Nums) == 0 {
//          panic("stack is empty")
//     }
//     return this.Nums[len(this.Nums)-1]
// }


// func (this *MinStack) GetMin() int {
//     if this.Head == nil {
//         panic("stack Head is nil")
//     }
//     return this.Head.Val
// }


// /**
//  * Your MinStack object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.Push(x);
//  * obj.Pop();
//  * param_3 := obj.Top();
//  * param_4 := obj.GetMin();
//  */
