func convert(s string, numRows int) string {
    //第一行的索引为 2*(n-1)*i  //n为传进来的那个numRows，i为第一行的第几位 0, 4, 8, 12
    //第二行的索引为 1,5,9...
    //第三行的索引为 2,6,10...
    //逐行拼接即可
    lenS := len(s)
    if lenS <=1 || numRows == 1 {
        return s
    }
    res := make([]byte,lenS)
    nowIndex := 0
    for i:=0; i<numRows; i=i+1{
        j:=0
        flag:=0 //当前是否进行到间隔位的字符
        for {
            if i == 0 || i == numRows-1 || flag == 0  {
                k := i+2*(numRows-1)*j
                if k < lenS{
                res[nowIndex] = s[k]
                nowIndex++
                j++
                flag=1
                } else {
                    break
                }
            } else {
                //中间位置的索引和前一个标准位置索引转换关系：
                //距离0个空格，下面的数字个数就是1，=》k+2
                //距离1个空格，下面的数字个数是1+2，=》k+2+2
                //距离n个空格，下面的数字个数是1+2*=》k+2+2*n
                //怎么算出有多少个空格? numRows-2为总空格数,当前第几行，空格数为 numRows-2-i
                k := i+2*(numRows-1)*(j-1) //这里k要重新计算，因为之前的k计算式中的j已经在上一轮++了，要还原
                if q:=k + 2 + (numRows-2-i)*2;q < lenS && q>=0{
					res[nowIndex] = s[q]
                nowIndex++
                } else {
                    break
                }
                flag = 0
            }
        }
    }
    
    return string(res)
}
