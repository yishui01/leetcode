//这题目是真TM的难炸天，这游戏能玩？？？
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    //中位数就是把一个集合分成两个相等长度的子集，有一个子集的值永远要大于等于另一个子集的值
    //只要满足下面两个条件，集合中的中位数就找到了
    // 1、len1 = len2  //左右两边的集合数量一样
    // 2、A[i-1] <= B[j] && B[j-1] <= A[i]   //左边的最大值反正要小于等于右边的最小值
    
    
    
    //通过设置变量i，通过i来推导出j的表达式，注意要从小的集合推导大数量的集合，反之会出现负数
    //二分法查找合适的i的位置
    //找到后根据总集合的长度判断中位数是直接返回还是两个子集的最小和最大值相加除以二
    
   
   
    //打算从短的集合开始定义变量，假设m是<n的，如果不是，调换位置,保证m始终<=n
    if (len(nums1) > len(nums2)) {
        tmp := nums1
        nums1 = nums2;
        nums2 = tmp
    }
    m, n := len(nums1) , len(nums2)  //两个集合的长度
    iMin := 0; iMax := m //设置i相关的变量
    
    //下面用二分查找找出合适的i的位置
    
    for iMax >= iMin{
       
        i := (iMax + iMin) /2
        //i出来了，那么j就可以根据i来推了
        //根据两边的集合数目相等可以得到 i + j = m-i + n-j
        //得出j = (m+n)/2 -i    i的取值范围为0~m 因此m必须要小于n，才能保证j永远不会出现负数（m就是那个小集合的长度）
        //有些人这里把j写为 (m+n+1)/2 - i ,多加了一个1，经过验证，两者写法影响的是当总长度为奇数的时候，中位数在左边还是右边
        //多加一个1代表j的值又多了1，代表分配给左边的数又多个一个（i和j代表分配的木板位置，i=0时左边集合数为0，i为1时左边
        //集合数为1，由此可以推出i为几左边的集合数量就为几，j也一样）
        j := (m+n)/2 - i 
        
        if i > iMin && nums1[i-1] > nums2[j] {
            //左边的a集合最大值大于右边B集合的最小值，说明i太靠右了，太大了，导致大的数都被分到左边了，要减小i
            //怎么减小，不是直接 i--，那个效率太低了，用二分法
            iMax = i-1
        } else if i<iMax && nums2[j-1] > nums1[i]{
            //j太大了，要减小，那就增加i的值就行
            iMin = i+1
        } else {
             maxLeft, minRight := 0,0
            
            if i == m {
                 //右边的为0
                minRight = nums2[j]
            } else if j == n {
                minRight = nums1[i]
            } else {
                if nums2[j] <= nums1[i] {
                minRight =  nums2[j]
                } else {
                     minRight =   nums1[i]
                }
            }
            //这要直接返回，不要把minRight和maxLeft都计算出来再去判断返回，那样在一个数组为空一个数组为1的时候，其中一边会溢出
             if (m+n)%2 == 1 {
                return float64(minRight)
            }
            
            //这里就是完美符合条件了  
            if i == 0 {
                //左边的为0
                maxLeft = nums2[j-1]
            } else if j == 0 {
                maxLeft = nums1[i-1]
            } else {
                if nums2[j-1] >= nums1[i-1] {
                    maxLeft =  nums2[j-1]
                } else {
                     maxLeft =  nums1[i-1]
                }
            }
            
            return float64(maxLeft+minRight)/2.0
            
        }
        
    }
    
    return 0.0
   
    
    
}