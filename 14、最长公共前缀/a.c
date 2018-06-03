char* longestCommonPrefix(char** strs, int strsSize) {
    //反正以第一个字符串为准，开始比对咯
   int i = 0, j =0,maxlen = 0;
   char curr = '\0';
   if(strsSize == 0)return "";
   if(strsSize == 1)return *strs;
    while(*(*strs+j)) {
        for(i = 1; i < strsSize; i++) {
            if(*(*(strs+i)+j) != *(*strs+j)){
                 *(*strs+maxlen) = '\0';
                 return *strs;
            }
        }
        maxlen++;
        j++;
    }
    *(*strs+maxlen) = '\0';
    return *strs;
}