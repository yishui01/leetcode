func largestRectangleArea(heights []int) int {
    //这题以为是dp，然后半天写不出一个正确的dp方程，无奈看网上答案，发现没有用dp的
    //解法1（耗时1400ms+，已废弃）：最大值图形总会以一个bar为高度，那么我们找每一个bar，由这个bar往两边扩散，看最多能扩散到多远,扩散到比自己矮的或者到边界就停下
    //解法2（采用的解法，耗时8ms）：从前往后遍历找到局部峰值点或者最后一个元素时，往前遍历，每次往前推一个的时候计算前推的bar和当前bar围成的面积，找出前推范围内的最大面积,就是当前bar的最大面积
    
    lens := len(heights)
    if lens == 0 {
        return 0
    }
    
    res := 0
    for i:=0; i < lens; i++ {
        if i+1 < lens && heights[i] <= heights[i+1]{
            continue
        }
        minH := heights[i]
        tmp := heights[i]
        for j:=i; j >=0; j-- {
            if heights[j] < minH {
                minH = heights[j]
            }
            nowArea := minH * (i-j + 1)
            if nowArea > tmp {
                tmp = nowArea
            }
        }
        
        if tmp > res {
            res = tmp
        }
    }
    
    return res
    
}
