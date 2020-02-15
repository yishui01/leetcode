func findDuplicates(nums []int) []int {
    //把对应下标的那个元素的值 * -1  ，如果有相同元素，那他们都会指向同一个下标，如果那个下标本身就是负数，那这个元素就是重复的
    res := make([]int,0)
    for _, v:= range nums{
        if v < 0 {
            v *= -1   //这里v是copy的，不会影响到原数组的计数，这里只是把v还原为原来的正值
        }
         if nums[v-1] < 0 {
                //之前就已经被打中过了
                res = append(res,v)
            } else {
                nums[v-1] *= -1 //打中
            }
      }

    return res
}


/*func findDuplicates(nums []int) []int {
     lens := len(nums)
     if lens <= 1 {
         return []int{}
     }
     //抽屉原理，1放到0号位，  2放到1号位、、、，最后那些重复的数字总是会占用着一个不属于他们的坑

    for i:=0;i<lens;i++{
        for nums[nums[i]-1] != nums[i] { //如果那个坑没有放置对应的萝卜，把萝卜移过去
            nums[nums[i]-1],nums[i] = nums[i],nums[nums[i]-1] 
        }
    }

    res:=make([]int,0)
    for k,v := range nums{
        if v != k+1 {
            res = append(res,v)
        }
    }

    return res
}*/
