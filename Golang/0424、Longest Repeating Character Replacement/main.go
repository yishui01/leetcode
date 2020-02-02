func characterReplacement(s string, k int) int {
    //滑动窗口On，还是不会做，这是一步步把窗口的size拉大，最后窗口的宽度就是最大连续字符串的长度
    //窗口宽度最开始是0，右边界反正是一直往右扩散，左边界呢，可以不动，也可以右移，左边界的不动的时候就是窗口被拉大的时候
    //那什么时候左边界会右移什么时候不动呢
    //这个窗口最大容量应该是  窗口内出现的最多的字符个数 + k    
    //那么一旦窗口超过了这个容量，那左边界就要右移了，右移的时候记得最左边的字符移出去

    left:=0
    maxTimes := 0 //历史窗口内出现最多的字符的个数

    /*
    执行用时12ms
    maps := make(map[byte]int,26)
    for right:=0;right<len(s);right++{
        maps[s[right]]++
        maxTimes = max(maxTimes,maps[s[right]]) //历史窗口内同一字符最大次数
        if right - left + 1 > maxTimes + k {
            //左边界左移
            maps[s[left]]--  //注意先把左边界的字符移出去，再移动左边界，不能先left++，不然left就不对了
            left++
        }
    }*/

    //将上面的maps转换为ints，从12ms降低到0ms，hash表的查找效率还是不行
    ints := make([]int,26)
    for right:=0;right<len(s);right++{
        ints[s[right]-'A']++
        maxTimes = max(maxTimes,ints[s[right]-'A']) //历史窗口内同一字符最大次数
        if right - left + 1 > maxTimes + k {
            //左边界左移
            ints[s[left]-'A']--  //注意先把左边界的字符移出去，再移动左边界，不能先left++，不然left就不对了
            left++
        }
    }


    return len(s)-left
}

func max(a,b int)int{
    if a > b {
        return a
    }
    return b
}
