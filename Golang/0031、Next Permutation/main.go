func nextPermutation(nums []int)  {
    //从后往前遍历，找到第一个数大于第二个数的点，就是要调换的点
    //调换过程为第二个数以后比他大的数放到原来第二个数的位置，再把原来的第二个数丢到后面，把后面的进行正序
    lens := len(nums)
    if lens == 0 {
        return 
    }
    isFind := false
    for i:=0; i < lens-1; i++ {
        if nums[lens-i-1] > nums[lens-i-2] {
            isFind = true
           
            point := nums[lens-i-2]
            pointindex := lens-i-2
            //找到point后面的数字里比point的大的那个最小值
            tmpMinMax := nums[lens - i - 1]
            tmpIndex := lens-i-1
          
            start := tmpIndex
            for j:=start; j < lens;j++ {
                if nums[j] < tmpMinMax && nums[j] > point {
                    tmpMinMax = nums[j]
                    tmpIndex = j
                }
            }
            //交换位置
            tmp := point
            nums[pointindex] = nums[tmpIndex]
            nums[tmpIndex] = tmp
            //再把剩下的正序排列
            sort.Ints(nums[pointindex+1:])
            break
        }
    }
    if !isFind {
        sort.Ints(nums)
    }
    return
}
