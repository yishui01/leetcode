type SummaryRanges struct {
    nums [][]int
}


/** Initialize your data structure here. */
func Constructor() SummaryRanges {
    return SummaryRanges{nums:[][]int{}}
}


func (this *SummaryRanges) AddNum(val int)  {
    //找头部节点有没有比当前val大1的，有的话合并
    //再找出尾节点有没有比当前节点小1的，有的话就合并，合并完之后再拿当前节点的尾部去和下一个元素的头部节点比对看是否能合并
    
    if len(this.nums) == 0 {
        this.nums = append(this.nums,[]int{val,val})
        return 
    }

    isMin := 0
    for i:=0;i<len(this.nums);i++{
         isMin = 0
        //先找头部节点有没有大1的
        if this.nums[i][0] == val {
            break
        }
        if this.nums[i][0] > val {
            //头结点已经比val大了
            if this.nums[i][0] == val +1 {
                //恰好可以合并
                this.nums[i][0] = val
                //此时不用考虑和前面区间合并的情况，因为前面区间的尾部如果比val小的话早就已经合并了，不会到这来
                //还有个就是头一个元素就比val大，然后到了这了，那也不用和前面的合并，因为是头一个元素，前面根本没有
            } else {
                //头结点大得val太多了，无法合并，直接添加到头节点之前吧
                head := make([][]int,i)
                copy(head, this.nums[:i])
                this.nums = append(append(head,[]int{val,val}),this.nums[i:]...)
            }
            break
        } else {
            //头节点小于val
            if this.nums[i][1] < val {
                //尾部节点也小于val
                if this.nums[i][1] == val-1 {
                    //刚好小1，合并，将尾节点替换为val
                    this.nums[i][1] = val
                    //再检查尾节点有没有和下一个元素的头结点达成合并关系，达成就合并
                    if i < len(this.nums)-1 && val == this.nums[i+1][0]-1{
                        this.nums[i][1] = this.nums[i+1][1]
                        this.nums = append(this.nums[:i+1],this.nums[i+2:]...)
                    }
                    break
                } else {
                    //小得太多了，直接continue
                    isMin = 1
                }
            } else {
                //尾部节点大于的话就直接break了,因为已经包含了
                break
            }
        }
    }

    if isMin == 1 {
        this.nums = append(this.nums,[]int{val,val})
    }
    
}


func (this *SummaryRanges) GetIntervals() [][]int {
    return this.nums
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
