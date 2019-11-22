func wiggleSort(nums []int)  {
    //一般先排序，分成两段，后半段往前半段插花还不行，一旦有相邻重复的元素那就G了
    //改进方案是分别从两段的末尾找元素一个个插花组合，奇数就把中位数算到前面那个数组中，这样就保证即使有相邻的重复元素也会被错开，因为两段是从末尾开始排序的
    //两个元素不是相邻元素，所以如果原数组合法（可以构成满足题意的数组）的话，那这两个元素一定不会重复

    sort.Ints(nums)
    res := make([]int,0)
    midIndex := (len(nums)+1)/2
    lens := len(nums)
    for i:=0;i<len(nums);i++{
       if i & 1 == 0 {
           midIndex--
           res = append(res,nums[midIndex])
       } else {
           lens--
           res = append(res,nums[lens])
       }
    }
    copy(nums,res)
}

