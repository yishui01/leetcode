bool isPalindrome(int x) {
    //把正序和倒序都存放到int数组里，两数组再逐个比对
    if(x < 0)return false;
    int len = 0, i = 0, copy = 0;
    copy = x;
    while(copy != 0) {
        copy/=10;
        len++;
    }
    if ( len == 0 || len == 1) return true;
    
    int asc[len],desc[len]; 
    copy = x;
    for(i = 0; i < len; i++) {
        desc[i] = asc[len-1-i] = copy%10;
        copy /= 10;
    }
    int res = true;
    for(i = 0; i<len; i++) {
        if(asc[i] != desc[i]){
            res = false;
            break;
        }
    }
    return res;
    
}