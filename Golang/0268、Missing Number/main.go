func missingNumber(nums []int) int {
  //解法1、找出理论上完美数组应有的总和，减去目前数组的总和，就是缺失的那个数
    lens := len(nums)
    sum := 0
    res := 0
    for i:=0;i<lens;i++{
        sum += nums[i]
        res += i
    }
     
    return res + lens - sum
}

//解法2、还可以将此数组与完美数组累计异或，异或到最后的结果就是那个缺失的数字

