func maxProduct(words []string) int {
    //感觉暴力无法AC，直接放弃，看答案，位运算，总共26个小写字母，一个int至少32位以上，足够了
    //判断两个单词是否重复就看相与是不是 == 0   == 0就代表有重复的，跳过
    
    if len(words) < 2 {
        return 0
    }
    ints := make([]int,0)
    
    for i:=0;i<len(words);i++{
        res := 0
        for j:=0;j<len(words[i]);j++{
            res = res | (1 << uint(words[i][j] - 'a'))
        }
        ints = append(ints,res)
    }
    // fmt.Println(ints)
    max,maxI,maxJ := 0,0,0
    for i:=0;i<len(ints)-1;i++{
        for j:=i+1;j<len(ints);j++{
             //fmt.Println("i=",i,"j=",j,"i&j=",(ints[i] & ints[j]),"lenMax=",len(words[i]) * len(words[j]) )
            if  ints[i] != 0 && ints[j] != 0 && (ints[i] & ints[j] == 0){
                if len(words[i]) * len(words[j]) > max {
                    max,maxI,maxJ = len(words[i]) * len(words[j]) ,i,j
                }
            }
        }
    }
    
    if max == 0 {
        return 0
    }
   
    return len(words[maxI]) * len(words[maxJ])
    
}
