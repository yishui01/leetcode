func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
    //直接用桶排序的思想，这种就用t作为桶的size，如果在同一个桶子里面的就不用说，肯定是符合条件的，直接返回true
    //如果没有，就把前一个桶子和后一个桶子的那个元素取出来，进行比对，如果绝对值<=t也是ok，此时还要利用滑动窗口
    //在遍历nums数组的时候要删除前面那些距离当前桶子的距离超过k的桶子，那种i和j的范围已经超过k了，不用保留了,不然比对的时候
    //从一个桶子中找到了元素，实际上这个桶子的元素是很久之前的了，桶子index已经超过k了，还美滋滋的比对成功，那就GG了
    
    if k <0 || t < 0 || len(nums) == 0 {
        return false
    }
    
    bucketSize := t+1 //t有可能为0 TMD，这里要+1，不然G了
    min := nums[0]
    
    for i:=1;i<len(nums);i++{
        if nums[i] < min {
            min = nums[i]
        }
    }
    
    buckets := make(map[int]int)
    
    for i:=0;i<len(nums);i++{
        bucketIndex := (nums[i]-min) / bucketSize
        if _,ok := buckets[bucketIndex];ok{
            return true //同一个桶子中居然有相同的元素，那k和t会直接满足
        }
        buckets[bucketIndex] = nums[i]
        
        if left,ok := buckets[bucketIndex-1];ok{
            if abs(left - nums[i]) <= t{
                return true
            }
        }
        
        if right,ok := buckets[bucketIndex+1];ok{
            if abs(right - nums[i]) <= t{
                return true
            }
        }
        
        //这一步别漏了，要删掉超过距离的元素
        if i >= k {
            delete(buckets,(nums[i-k]-min)/bucketSize)
        }
        
    }
    
    return false
    
    
}
func abs(a int) int {
    if a<0{
        return -a
    }
    return a
}


// func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
//     //暴力流，O(n2)勉强通过，208ms
//     for i:=0;i<len(nums)-1;i++{
//         for j:=i+1;j<len(nums);j++{
//             if abs(nums[i]-nums[j]) <= t && abs(i-j) <= k {
//                 return true
//             }
//         }
//     }
    
//     return false
// }

// func abs (a int) int {
//     if a<0{
//         return -a
//     }
//     return a
// }
