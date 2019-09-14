func largestNumber(nums []int) string {
    //本人之前的思路（注意是错的，是错的，是错的！）：数字全部转换成string，string好像直接比较的话是逐位比对，正好符合要求，找出最大的即可，这里会挂在某些case比如3,30 判断出30比3大，因为30比3多了一位，所以会排序成303,而实际上正确的排序是330...
    
    //大佬思路（正确）：这种解法对于两个数字a和b来说，如果将其都转为字符串，如果ab > ba，则a排在前面，比如9和34，由于934>349，所以9排在前面，再比如说30和3，由于303<330，所以3排在30的前面。按照这种规则对原数组进行排序后，将每个数字转化为字符串再连接起来就是最终结果
    lens := len(nums)
    if lens == 0 {
        return ""
    }
    var s mySlice
    for i:=0;i<lens;i++{
        s = append(s,strconv.Itoa(nums[i]))
    }
    sort.Sort(s)
    if s[0] == "0"{
        return "0"
    }
    str := ""
    for i:=0;i < len(s);i++{
        str += s[i]
    }
    return str
}

type mySlice []string

func (s mySlice)Len() int {
    return len(s)
}

func (s mySlice)Less(i,j int) bool {
    return s[i]+s[j] > s[j]+s[i]
}

func (s mySlice)Swap(i,j int) {
    s[i],s[j] = s[j],s[i]
}


