func sortColors(nums []int)  {
    //左右指针吧，左边指针指着当前第几个0，右边指针指着当前第几个2
    //遇到0就按照左指针位置+1存入（交换），遇到1就跳过，不管，因为1在0和2排完之后自动就到中间了，遇到2就右指针-1存入（交换）
    //坑1、遇到0可以直接交换，把0前换，因为被交换者是被遍历过的，如果遇到2的话就不能直接交换，因为把2往后换，被交换者还没被遍历过呢，然后就被换到前面去了
    //然后i如果继续+1的话，那遍历就到后面去了，这个被交换者就被遗漏了，所以不是所有的都要i++,一旦遇到2，交换者交换后，被交换者如果是0或者2这种特殊数字，就要停止i++
    //坑2、right指针一直在左移，如果，循环的i一直在增大，那么遍历到那些已经交换好了的2的时候又会和已经很小的right指针交换，然后又把2换到前面去了
    //所以一旦right指针到了遍历的那个i相等，就break，排序完成
    lens := len(nums)
    if lens == 0 {
        return 
    }
    
    left := -1
    right := lens
    
    for i:=0;i < lens; { //注意这里没有i++
        if i>=right {
            break
        }
        if nums[i] == 0 {
            left++
            tmp:=nums[left]
            nums[left] = nums[i]
            nums[i] = tmp
            i++
        } else if nums[i] == 2{
            right--
            tmp:=nums[right]
            nums[right] = nums[i]
            nums[i] = tmp
            if nums[i] != 0 && nums[i] != 2 {
                i++
            }
        } else {
            i++
        }
        
    }
    
    

}
