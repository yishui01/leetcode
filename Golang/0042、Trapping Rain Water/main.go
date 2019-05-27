func trap(height []int) int {
    //看到一个很强的思路：整个图形在被填充满之后是一个塔形
    //从左往右是先升高后降低，那么我们只要找到塔顶，然后塔顶两边都是相对较低的台阶，我们找到各个元素与所对应的台阶高度的差值即可
    
    lens := len(height)
    if lens <=2 {
        return 0
    }
    
    //找到塔顶
    topHeight := -1
    topIndex := -1
    for i:=0; i < lens; i++ {
        if height[i] > topHeight {
            topHeight = height[i]
            topIndex = i
        }
    }
    area := 0
    nowTop := height[0]
    //计算左边的面积
    for i:=0; i < topIndex; i++ {
        if height[i] > nowTop {
            nowTop = height[i]
        } else {
            area += nowTop - height[i]
        }
    }
    //右边的面积
    nowTop = height[lens-1]
    for i:=lens-1; i >topIndex; i-- {
        if height[i] > nowTop {
            nowTop = height[i]
        } else {
            area += nowTop - height[i]
        }
    }
    
    return area
    
}
