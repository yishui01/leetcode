/*
给定一个字符串，找出不含有重复字符的最长子串的长度。

示例：

给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。
**/

/**
这题目解决思路是遍历字符串，将不重复的字符保存到数组中，每次获取下一个字符时，遍历数组
看是否已经存在，存在的话，敲黑板！！！（将数组中重复的字符的索引后的字符全部向前推进一位，再把新来的字符加到尾部）。

比如abcabcbb，前3次已经将abc加到了数组中，现在又来了个a，将数组中重复的字符索引=》a的索引，后面的字符全部向前
推进一位=》a后面的字符是b,c，全部向前推进，此时数组变为了b,c,c，再把新来的字符加到尾部=》再把新来的a加到b,c的后面，覆盖c，
此时数组变为了b,c,a，再下一次循环
**/
int lengthOfLongestSubstring(char* s) {
    if(!*s)return 0;
    int max_len = 0;
    int length = 0;
    int i = 0, j = 0, final, si, z; // j记录着当前数组有多少个元素
    length = strlen(s);
    char temp[length];
    while (*s) {
        int flag = 0; //数组中是否存在
        //遍历数组检查
        for (i = 0; i<j; i++) {
            if (temp[i] == *s) {
                flag = 1;
                break;
           }
        }
        if (flag == 0) {
           
             temp[j] = *s;    
             j++; 
        } else {
            if (max_len < j) {
                max_len = j; 
            }
          z = temp[i];
          final = 0;
          si = i+1;
          for(final = 0; final<j-i-1; final++){
			  //整体前移，重新填充数组
              temp[final] = temp[si++];
          }
         temp[final] = z; //将重复的字符加到尾部
         j = j-i;
        }
       
        if (max_len < j) {
          max_len = j; 
        }
       
        s++;
    }
   
    return max_len;
    
}