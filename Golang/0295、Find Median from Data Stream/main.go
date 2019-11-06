type MedianFinder struct {
    nums []int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{}
}


func (this *MedianFinder) AddNum(num int)  {
    //直接插入排序
    //从头往后遍历，直到找到一个比自己大的数，添加到他前面就行了
    start:=len(this.nums)
    for i:=0;i<len(this.nums);i++{
        if this.nums[i] >= num {
            start = i
            break
        }
    }
    
    newNums := make([]int,0)
    newNums = append(newNums,this.nums[:start]...)
    newNums = append(newNums,num)
    newNums = append(newNums,this.nums[start:]...)
    this.nums = newNums
}


func (this *MedianFinder) FindMedian() float64 {
    if len(this.nums) == 0 {
        return 0
    }

    if len(this.nums) % 2 == 0 {
        return (float64(this.nums[len(this.nums)/2]) + float64(this.nums[len(this.nums)/2-1])) / 2
    }
    return float64(this.nums[len(this.nums)/2])
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */



