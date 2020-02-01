func originalDigits(s string) string {
    /**
    “z” 只在 “zero” 中出现。 0
    “w” 只在 “two” 中出现。  2
    “u” 只在 “four” 中出现。 4
    “x” 只在 “six” 中出现。  6
    “g” 只在 “eight” 中出现。8

    “h” 只在 “three” 和 “eight” 中出现。 3 - 8
    “f” 只在 “five” 和 “four” 中出现。   5
    “s” 只在 “seven” 和 “six” 中出现。   7

    “i” 在 “nine”，“five”，“six” 和 “eight” 中出现。 9
    “n” 在 “one”，“seven” 和 “nine” 中出现。 1
    **/

    maps := make(map[rune]int,len(s))
   for _,v := range s {
       maps[v]++
   }

   var ints [10]int   //代表0-9每个数字的个数

   ints[0] = maps['z']
   ints[2] = maps['w']
   ints[4] = maps['u']
   ints[6] = maps['x']
   ints[8] = maps['g']

   ints[3] = maps['h']-ints[8] //3
   ints[5] = maps['f']-ints[4] //5
   ints[7] = maps['s']-ints[6] //7

    ints[9] = maps['i']-ints[5]-ints[6]-ints[8]
    ints[1] = maps['n']-ints[7]-ints[9]*2 //这里要注意，一个9占了两个n所以要乘以2
    var bs bytes.Buffer  //循环拼接字符串使用buffer优化  //76ms =》 4ms
	for k,v := range ints {
        tmpk := byte(k)+'0'
		for i:=0;i< v;i++ {
			bs.WriteByte(tmpk)
		}
	}
	return bs.String()
}
