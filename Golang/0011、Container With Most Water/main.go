func maxArea(height []int) int {
   //最简单的是遍历
    //第二种是双指针法，双指针就是一头一尾两个指针，移动矮的那个往中间靠拢，因为面积的高度是由矮的决定的，
    //宽度反正一直在缩小，所以要尽可能提升矮子的高度。为什么不会漏掉最优解？
    //贪心算法，你的每一次缩短宽度，都是选到了目前最优的移动方案，并且每一次的移动都是一格，是最小颗粒度，最终能达到最优解
    //用数学归纳法证明，通过每一步贪心选择，最终可得到问题的一个整体最优解
    
    /**
    假设：该算法并没有遍历到容量最大的情况
    我们令容量最大时的指针为p_left和p_right。根据题设，我们可以假设遍历时左指针先到达p_left，但是当左指针为p_left时，右指针还没有经过p_right左指针就移动了
    已知当左指针停留在p_left时，它只有在两种场景下会发生改变

    左指针和右指针在p_left相遇，则右指针一定在前往p_left的途中经过p_right，与题设矛盾
    右指针位于p_right右侧且当前的值大于左指针。则在这种情况下，此时容器的盛水量为（比最大面积还宽的宽度*p_left高度）比题设中最大的盛水量还要大，与题设矛盾（此时不用管右侧指针和p_right谁高，右侧指针就算小于p_right又怎样，此时的右侧指针造成p_left移动的话，那右侧肯定>p_left,   高度还是由矮的p_left决定）
    */
    maxArea,pLeft , pRight := 0, 0, len(height)-1
    
    for {
        if pRight == pLeft {
            break
        }
        minheight := 0
        if (height[pRight] >= height[pLeft]) {
             minheight = height[pLeft]
             pLeft++ 
        } else {
             minheight = height[pRight]
             pRight--
        }
        nowArea := (pRight - pLeft+1) * minheight
        if maxArea < nowArea {
            maxArea = nowArea
        }
    }
    
    return maxArea
}