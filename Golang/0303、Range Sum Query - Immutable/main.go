type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    tmps := make([]int,len(nums))
    if len(nums) > 0 {
        tmps[0] = nums[0]
    }
    for i:=1;i<len(tmps);i++{
        tmps[i] = tmps[i-1]+nums[i]
    }
    return NumArray{nums:tmps}
}


func (this *NumArray) SumRange(i int, j int) int {
    if j == 0 || i == 0 {
        return this.nums[j]
    }
    
   return this.nums[j]-this.nums[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
