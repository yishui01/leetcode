func splitArray(nums []int, m int) int {
    //二分查找，跪了，二分还能这么用的，反向推理
    //要分成m个数组，假设m = 1 ,那res就是所有元素之和sum(nums),如果m = len(nums), 那 res 就是 max(nums)
    //设l = max(nums),r = sum(nums)  求出mid ，每次来一个新mid时，遍历数组，看有多少个子数组和大于该mid
    //多于m说明mid小了，随便来几个元素就能组成一个合法子数组，但其实合法数组没有那么多
    //小于m说明mid大了
    
    l,r := 0,0
    for _,v := range nums {
        if v > l {
            l = v
        }
        r += v
    }

    for l < r {
        mid := l + (r-l) >> 1
        validCount := 1  //初始化为1，代表当前子数组的个数，一开始都没有拆，整个就是一个大数组，所以子数组个数就是这个大数组，就是一个
        tmpSum := 0
        for _,v := range nums {
            tmpSum +=v
            if tmpSum > mid {
                validCount++ //拆出一个合法子数组，validCount++
                tmpSum = v
            }
        }
        if validCount > m { //mid小了
            l = mid+1
        } else {  //mid太大或者mid刚好合适，都要缩小mid，因为mid即便刚好能分出m组，此时mid代表子数组的和里面的max，那题意是要找出最小的max
                    //所以尝试缩小下mid，看能不能还能分出m个子数组，这样一直不断的缩小mid即可，缩到l等于r，那就退出循环
            r = mid
        }
    }
    //返回r是由于mid已经缩小的不能再缩了，此时mid就是那个最小的合法的子数组和。
    //这里返回l或者r都可以，因为l和r最后肯定都是相等的，因为l迭代和r迭代只能是mid或者mid+1，只会在l和r的区间，不可能超出这个区间
    return r
    //为什么mid是最小数组和？反证法：
    //假设mid比正确答案（最小切割序列）的最大子数组和还要小，设该正确答案最大子数组和为x, 即 mid < x,
    //此时用mid切割出来的数组序列中，所有子数组的和都<=mid ，又由于mid < x,因此所有子数组和也都 < x，导致正确答案并非最小切割序列，自相矛盾
    //假设mid > x，那么此时必然存在另一个m2，m2 < mid,并能切出m个数组，但由于此时m已经无法再小，再小就切不出m个了，所以不存在m2

    //还有个思路证明mid：
    //首先正确答案（最小切割序列的最大子数组和）肯定在l和r之间，那么二分法时，l的迭代只会排除掉不合法的数字，r的迭代会保留合法mid，直到更小的合法mid迭代掉r
    //而 res 就是那个最小的我们要找的mid，所以mid最终会被找到，并赋值给r
  
}
