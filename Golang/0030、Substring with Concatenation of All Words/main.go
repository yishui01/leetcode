func findSubstring(s string, words []string) []int {
    // 滑动窗口模式，这个题目有个隐藏的点就是所有的正确的解，他们之间的偏移距离不会超过一个单词的长度。
    // 比如说单词是 aba、bab   那么在字符串中，要找到所有的解，只需要将滑动窗口的起始位置分别放在
    //0号、1号、2号索引（因为单词长度是3，所以放三个位置），
    //然后一个个的加单词即可  假设字符串为abababab ，滑动窗口起始位置分别放在0,1,2时，得出的答案也分别为0,1,2
    //这样就能在答案有重叠的情况下全部找出,并且能够套出所有正确答案，就是不管答案前有多少个干扰字符
    //只要滑动窗口的起始位置涵盖了一个单词的长度，那总有一个窗口一路套下来，能把答案套上
    // 比如 单词时 skt、ssb
    //1、输入的字符串为 sktssb 0号位窗口直接套中
    //2、输入字符串为 asktssb  1号位窗口套中
    //3、输入字符串为 aasktssb 2号位窗口套中
    //4、输入字符串为 aaaasktssbn 1号位窗口套中
    //...  反正总有一个窗口能把答案给套出来
    
    arrLen := len(words)
    sLen := len(s)
    if arrLen == 0 || sLen == 0 ||  len(words[0])*arrLen > len(s) {
        return []int{}
    }
    wLen := len(words[0])
    subLen := wLen*arrLen
    overLen := sLen - subLen //左边框界限，不能超过这个index，超过了那剩余字符串就小于答案字符串了，不可能存在答案
    var count int //当前合法的单词个数，用于判断是否已经达成正确解
    res := make([]int,0)
    
    remain := make(map[string]int) //这个map是保存合法word还能出现的剩余次数
    left, right := 0, 0
    
    
    // moveLeft 让 left 指向下一个单词
	moveLeft := func() {
		// left 指向下一个单词前，需要修改统计记录
		// 增加 left 指向的 word 可出现次数一次，
		remain[s[left:left+wLen]]++
		// 连续符合条件的单词数减少一个
		count--
		// left 后移一个 word 的长度
		left += wLen
	}
    
    tmpRemain := make(map[string]int)
    for _,v := range words {
        tmpRemain[v]++
     }
    
    //这里要借助一个中间变量tmpRemain来帮助重置
    count = 1  //设置为1是为了避免下面的第一次reset未初始化
    reset := func(){
        if count == 0 {
            //当count为0时，不需要重置map，此时为了避免第一次count的初始值为0，导致没有初始化，
            //那么就在上面将count的初始值设置为1，真的细节 执行时间从48ms优化到4ms 6的一批
            return
        }
        for k,v := range tmpRemain {
            remain[k] = v
        }
        count = 0
    }
    
    reset()
    
    //滑动窗口的所有起始位置
    for i:=0;i < wLen;i++ {
        left,right = i,i
        reset()
        //下面这个循环是开始拖同窗口的左右边框，每次拖动的距离固定为一个单词长度（因为所有单词都是定长）
        for left <= overLen{
            //剩余字符串长度不能小于正确答案字符串长度
            word := s[right:right + wLen] //新word
            remainTimes, ok := remain[word]
            switch {
                case !ok:
                // word 不在 words 中
				// 从 right+lenw 处，作为新窗口，重新开始统计
                left, right = right+wLen,right+wLen
                reset()
                case remainTimes == 0:
                //单词出现次数过多,往右边拉动左边框，拉动的时候要重置原来最左边的单词的出现次数
                moveLeft()
                default:
                //ok && remainTimes == 0，word 符合出现的条件
                // moveRight
                remain[word]--
                count++
                right += wLen //拉动右边框
                if count == arrLen {
                    //达成最优解
                    res = append(res, left)
                    //移动左边框
                    moveLeft()
                }
                
            }
        }
    } 
    
    return res
    
    
}
