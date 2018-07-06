char* longestPalindrome(char* s) {
    /*每个字符向两边扩散，回文串长度有可能是奇数，也有可能是偶数，所以每次扩散时都要考虑奇偶，奇偶的区别在于起点不同，其他一样*/
    int len = strlen(s);
    if (len == 1) return s;
    int maxlen = 0, currlen, i, left, right, sbegin = 0;
    for (i = 0; i<len; i++) {
        //考虑回文串是奇数的情况
        left = i - 1; right = i + 1;
        while(left>=0 && right<len && (s[left] == s[right])) {
            currlen = right - left;
            if (maxlen < currlen) {
               
                maxlen = currlen;
                sbegin = left;
            }
            left --;  right ++; //扩散
        }
        
        //考虑回文是偶数的情况
        left = i; right = i+1;
        while (left >= 0 && right < len && (s[left] == s[right])) { 
            currlen = right - left;
            if (maxlen < currlen) {
                maxlen = currlen;
                sbegin = left;
            }
            left--; right++;
        }
    }
    
    char* res = (char *) malloc(sizeof(char) * (maxlen+25));
    strncpy(res, s+sbegin, maxlen+1);
	res[maxlen+1] = '\0';
    return res;
    
}