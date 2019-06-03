func jump(nums []int) int {
    //看了大佬的解法，我又要怀疑人生了，我究竟适不适合写代码？o(╥﹏╥)o
    //这个题有个提示，就是无论你怎么跳，始终是有解的
    //那我们就找到最优最远的跳法就行
    
    //大佬的解法是这样，找到当前点的最远范围，只要i还在这个范围内，那就使劲找，找什么呢，
    //找到那个最大的数，这个最大的数可以帮你把最远距离撑开，撑得更远，撑得越远越好，如果撑到头了，那答案也就有了
    last := 0 //当前查找的返回
    curr := 0 //当前能撑到的最远距离
    res := 0//步数
    lens := len(nums)
    if lens <=1 {
        return 0
    }
    for i:=0; i < lens; i++ {
        if i > last {
            //如果i超过了可以查找的范围，那就要重新更新查找范围，就是向前跳一步
            res++
            last = curr
        }
        if curr < i+nums[i] {
             curr = i+nums[i]
        }
        //下面这个判断是我优化的，就问你怕不怕，哈哈哈
        //这个已经撑到最远距离了，那答案也就有了
        //每一次生成curr的时候都会有这个判断，所以不会出现curr还在last中的情况，因为那样的话，那个last还是curr的时候早就return了
        if curr>= lens-1 {
            return res+1
        }
    }
    
    return res
    
}
// func jump(nums []int) int {
//     //先用回溯试一试，好吧超时，OJ大佬手下留点情啊
//     lens := len(nums)
//     if lens == 0 {
//         return 0
//     }
//     if lens ==1 {
//         return 0
//     }
//     if lens <=2 {
//         return 1
//     }
    
//     //把每一种跳法所需要的步数都记录到数组中
//     res := -1
    
//     jumpRecursion(nums, 0, 0, 0,lens-1, &res)
    
//     return res
    
    
// }

// /**
// tmpRes 上一层的结果
// nowIndex 这一层要累加的索引
// step 上一层一共使用的步数
// target 目标
// res 结果数组
// **/
// func jumpRecursion(nums []int,  tmpRes int, nowIndex int, step int, target int, res *int) {
//     if nowIndex > target {
//         return 
//     }
//     if tmpRes ==  target {
//         if step < *res || *res == -1 {
//             *res = step
//         }
//     }
    
//     for i:=0; i < nums[nowIndex]; i++ {
//         //跳下一层
//         jumpRecursion(nums, tmpRes + i+1, nowIndex+i+1, step + 1, target, res);
//     }
    
// }
