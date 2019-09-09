func twoSum(numbers []int, target int) []int {
   //由于题目说必定只有一个答案，并且数组数排过序的，那就一头一尾直接相加就行，大了就右指针左移，小了就左指针右移
    lens := len(numbers)
    left, right := 0, lens-1
    
    for left < right {
        tmp := numbers[left] + numbers[right]
        if tmp == target {
            return []int{left+1,right+1}
        } else if tmp > target {
               right--
        } else {
            left++
        }
    }
    
    return []int{left,right}
   
}


// func twoSum(numbers []int, target int) []int {
//     //直接O(n),利用maps存储之前出现过的数字
//     maps := make(map[int]int)
//     lens := len(numbers)
    
//     for i:=0;i<lens;i++ {
//         if val,ok := maps[target-numbers[i]];ok {
//             return []int{val,i+1}
//         }
//         maps[numbers[i]] = i+1
//     }
    
//     return []int{}
// }
