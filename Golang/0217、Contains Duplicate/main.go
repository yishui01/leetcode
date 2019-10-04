func containsDuplicate(nums []int) bool {
    //直接用map
    maps := make(map[int]bool)
    for i:=0;i<len(nums);i++{
        if _,ok := maps[nums[i]];ok {
            return true
        }
        maps[nums[i]] = true
    }
    return false
}
