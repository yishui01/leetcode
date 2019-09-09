func majorityElement(nums []int) int {
    //摩尔投票，只有在确定存在过半数时才会生效，跪了，O(n)时间 O(1)空间,摩尔警告。。泰拳警告，尼玛这也太叼了啊，我fou了
    lens := len(nums)
    res := nums[0]
    count := 0
    for i:=0;i<lens;i++ {
        if count == 0 {
            res = nums[i]
        }
        if nums[i] == res{
            count++
        } else {
            count--
        }
    }
    return res
}

// func majorityElement(nums []int) int {
//     lens := len(nums)
//     times := lens/2
//     maps := make(map[int]int)
//     for i:=0;i<lens;i++ {
//         maps[nums[i]]++
//         if maps[nums[i]] > times {
//              return nums[i]
//         }
//     }
//     return 0
// }
