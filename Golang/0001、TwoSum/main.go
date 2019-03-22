//之前用的双重循环，这次借鉴了discuss中的算法，很巧妙的算法
func twoSum(nums []int, target int) []int {
     all := make(map[int]int, len(nums))
        for i:=0; i < len(nums); i++ {
            value, exists := all[nums[i]]
                if exists {
                    return []int{value, i}
                } else {
                    all[target - nums[i]] = i
                }
        }
   
    return []int{}
}
