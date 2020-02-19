func findDisappearedNumbers(nums []int) []int {
    //还是用之前的抽屉原理，桶排序
    for i:=0;i<len(nums);i++{
        for nums[nums[i]-1] != nums[i]{ //如果在他该有的位置上不是这个数字，替换，假设nums[2]=5  实际上 nums[5-1]就应该等于5才对，不等于就进入for
             nums[nums[i]-1],nums[i] = nums[i],nums[nums[i]-1]
        }
    }
    res := []int{}
    for k,v := range nums {
        if v != k+1 {
            res = append(res,k+1)
        }
    }

    return res
}
