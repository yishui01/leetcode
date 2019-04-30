func threeSum(nums []int) [][]int {
    //把三个数的和化成两个数的和,一个一个来，先找一个，再以这个找到的为准，从剩下的数字中找两数之和=第一个数的负数的
 
    allMap := make(map[string][]int)
    res := make([][]int,0)
    lens := len(nums)
    sort.Ints(nums)
    if lens == 0 {
        return [][]int{}
    }
    if nums[0] == nums[lens-1] && nums[0] == 0 && lens >= 3{
        a := make([][]int,0)
        return append(a,[]int{0,0,0})
    }
    
    
    for i:=0; i < lens; i++ {
        tmp := nums[i]
        tmpMap := make(map[int]int)
        //从接下来的数字中找两数之和为-tmp的
        for j:=0;j <lens;j++ {
            if j == i {
                continue
            }
            if _,ok := tmpMap[nums[j]];ok{
                //找到了,添加到结果map中, key为各值排序后的字符串,value为
                allMap[implode([]int{nums[i],nums[j],nums[tmpMap[nums[j]]]})] =  []int{nums[i],nums[j],nums[tmpMap[nums[j]]]}
            } else {
                tmpMap[-tmp - nums[j]] = j
            }
        }
    }
    i:=0
   // fmt.Println(allMap)
     for _, value := range allMap {
        res = append(res,value)
        i++
    }
    
    return res
}

func implode(s []int) string{
    sort.Ints(s)
    return strconv.Itoa(s[0])+strconv.Itoa(s[1])+strconv.Itoa(s[2])
}



