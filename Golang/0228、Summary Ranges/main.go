func summaryRanges(nums []int) []string {
    lens := len(nums)
    res := make([]string,0)
    if lens == 0 {
        return res
    }
    
    start := nums[0]
    end := nums[0]
    
    for i:=1;i<lens;i++{
        if nums[i]!=end+1{
            //截断
            if start == end {
                res = append(res, strconv.Itoa(start))
            } else {
                res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(end))
            }
            start = nums[i]
            end = nums[i]
        } else {
            end = nums[i]
        }
    }
    
    //把最后的那部分也加进去，因为之前的触发机制是在遇到截断区间时才触发append
     if start == end {
          res = append(res, strconv.Itoa(start))
       } else {
          res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(end))
      }
    
    
    return res
}
