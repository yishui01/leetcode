type Solution struct {
    nums []int
}


func Constructor(nums []int) Solution {
    rand.Seed(time.Now().UnixNano())
    return Solution{nums}
}


func (this *Solution) Pick(target int) int {
    //可以直接用hashmap、、、想了下还是用蓄水池抽样吧，给点面子╮(╯▽╰)╭
    res := -1
    count := 1
    for k,v := range this.nums{
        if v == target{
            if rand.Intn(count) < 1 { //注意这里只能为小于，不能为小于等于，假设有三个数，randIntn是前闭后开，所以rand(1)  rand(2)  rand(3)最能产生小于1、小于2、小于3的数
            //如果你是<=1  那么到第二个数时，依然为百分百，因为返回的整数只能为0和1，而这百分百符合判断条件，错误，应设为 < 1
                //打中蓄水池，OK
                res = k
            }
            count++
        }
    }

    return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
