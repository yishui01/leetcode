type Solution struct {
    origin []int
    nums []int
}


func Constructor(nums []int) Solution {
    rand.Seed(time.Now().Unix())
	tmp1 := make([]int, len(nums))
	tmp2 := make([]int, len(nums))
	copy(tmp1, nums)
	copy(tmp2, nums)
    return Solution{
    	origin: tmp1, 
    	nums: tmp2,
    }
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    copy(this.nums, this.origin)
    return this.nums
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    for i:=0;i<len(this.nums);i++{
        randomIndex := rand.Intn(len(this.nums)-i)+i
        this.nums[i],this.nums[randomIndex] = this.nums[randomIndex],this.nums[i]
    }
	
    return this.nums
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
