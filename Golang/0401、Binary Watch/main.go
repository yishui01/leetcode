func readBinaryWatch(num int) []string {
    //由于只有12个小时数、60个分钟数，两层循环，小时数和分钟数的二进制1的个数符合要求，既是可以表示的时间
    
    res := make([]string,0)
    for i:=0;i<12;i++{
        for j:=0;j < 60;j++{
            h := fmt.Sprintf("%b",i)
            m := fmt.Sprintf("%b",j)
            countOne := strings.Count(h,"1")+strings.Count(m,"1")
            if countOne == num{
                res = append(res, fmt.Sprintf("%d:%02d", i, j))
            }
        }
    }
   

    return res

}
