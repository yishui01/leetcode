char* convert(char* s, int numRows) {
    //一行一行来，Z字形的每个整列都是等差数列，整列与整列之间的单列也和整列有等差关系
    int i,j, diff, len, currenIndex; //diff为每个整列之间的差值
    len = strlen(s);
    if(len <= numRows || numRows == 1) return s;
    char * res = (char *)malloc(sizeof(char) * strlen(s)+1);
    j = 0;
    diff = 2 * numRows -2; //这等式由观察法得来
    for (i =0; i < numRows; i++) {
        currenIndex = i;
        while (currenIndex< len) {
                   //先把整列的加了
                    res[j] = s[currenIndex];
                    j++;
                    if(i > 0 && i < numRows - 1) {
                        //如果不是首行或者尾行，把与下一个整列之间的单列也加了
						//currenIndex是本列本行字符的索引，加上diff就是下一个整列的同行字符索引
						//行数越大，单列字符的索引离下一个整列的同行字符索引越远，且与行数成2的倍数
                        if(currenIndex + diff - 2 * i < len) {
                            res[j] = s[currenIndex + diff - 2 * i];
                            j++;
                        }
                    }
            currenIndex += diff;
        }

    }
    res[len] = '\0';
    return res;
}