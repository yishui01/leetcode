func containsNearbyDuplicate(nums []int, k int) bool {
    //这...还是用个map存下?
    maps := make(map[int]int)
    for i:=0;i<len(nums);i++{
        if val,ok := maps[nums[i]];ok{
            if i-val <= k {
                return true
            }
        }
        maps[nums[i]] = i
    }
    
    return false
}
