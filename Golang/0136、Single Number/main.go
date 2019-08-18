func singleNumber(nums []int) int {
    //之前想到了位图，但是一想到位操作有点麻烦就没继续往下想了，实际上主要是不熟悉位操作，位操作里有个异或，可以直接解答本题
    res :=0
    for _,v := range nums {
        res ^= v
    }
    
    return res
}
