func findTheLongestSubstring(s string) int {
    //为啥我想到了前缀和，但是还是不知道做？
    //具体题解: https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/solution/jiang-ti-mu-yi-bu-bu-deng-jie-zhuan-hua-fei-qi-ji-/
    //好吧思路如下：
    //求[i,j]区间范围内数字出现偶数次
    //转换为[0,i-1]和[0,j]的状态相同，想想前缀和
    //为什么这么转化？假设[0,i-1]此时子串满足都为偶数个，然后再继续往后，到了[0,j] 这时候又变为了全偶数，那么期间必定经历了0次转换或者2次或者4次，注意0次也是偶数，
    //考虑两次的情况，从最开始偶变奇（出现次数+1），再从奇变偶（出现次数+1），因此一共出现2次，满足题意
    //再来看奇-奇，一样的，奇-偶 一次，偶-变回奇一次，也是经历了两次，那代表这个区间这个数字也是出现了两次，满足
    //举例 el 这个单词，从0号位e开始，e是元音，将第二位二进制置为1，tmp变为2，并记录此时索引值0，赋值给val，最终映射为2=》0，再遍历下一个字符，el，l不是元音字母，因此tmp还是2
    //然后之前maps中已经有个2了，也就是找到了一个值相等的区间，满足上面条件，因此最大长度为el中的l字符，代表aeiou各出现了0次，因此最长子串是1，res = k-val = 当前索引-第一次出现的索引 = 1-0 = 1
    //再来，考虑bcbcbc，这里要注意的地方就是，可能某个子串或者整个字符串就都是元音字母出现偶数次的情况，那么此时根据异或规则，tmp值整体为0，那，res=k-val ,而此时的val应该
    //从maps中找出，找maps[0]的值，并计算res，推出此时0的那个映射值应该为-1，这样才能满足res=k-(-1) = k+1 = 当前字符串的长度，所以要设置maps[0]=-1
    
    //初始化一map记录各种状态最小的下标，5个元音字母，每个字母的奇偶只需要1位表示，总共是2的5次方，一共是32种状态

    res, tmp := 0, 0
    maps := make(map[int]int,32)
    maps[0]=-1
    for k,v := range s {
        switch v {
            case 'a':tmp ^=1
            case 'e':tmp ^=2
            case 'i':tmp ^=4
            case 'o':tmp ^=8
            case 'u':tmp ^=16
        }
        if val,ok := maps[tmp];!ok{
            maps[tmp] = k
        } else{
            res = max(res,k-val)
        }
    }
    return res

}

func max(i,j int)int{
    if i>=j{
        return i
    }
    return j
}
