/*
这题目的执行时间也太恐怖了，之前最外层while循环的条件是*haystack，执行时间是500ms左右，
然后在最外层while循环中加了长度判断，发现执行时间变成了4ms，ememmmmmmmm，我该怎么说才好，总之是个很有用的经验,
当总字符串的剩余长度小于子字符串的时候，其实就不用遍历了，直接返回-1
*/
int strStr(char* haystack, char* needle) {
    if(!*needle)return 0;
    char * tmpneed, *tmpall;
    int index = 0, flag = 0, alllen = strlen(haystack), needlen = strlen(needle);
    while(*haystack && alllen >=needlen) {
        if(*haystack == *needle) {
            tmpneed = needle;
            flag = 0;
            tmpall = haystack;
            while (*tmpneed) {
                if(*tmpneed != *tmpall) {
                    flag = 1;
                    break;
                }
                tmpneed++;
                tmpall++;
            }
            if(flag == 0)return index;
        }
        index++;
        haystack++;
        alllen--;
    }
    
    return -1;
}