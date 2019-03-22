int reverse(int x) {
   
    int i = 0, res = 0, len = 0, copy;
    copy = x;
    while (copy != 0) {
        copy /= 10;
        len++;
    }
    
    while(x != 0) {
        if (x > 0) {
             if(res + x % 10 * pow(10, len-i-1) > pow(2, 31) -1){
                 return 0;
             }
        } else {
             if(res + x % 10 * pow(10, len-i-1) < pow(-2, 31)){
             return 0;
         }
        }
        res += x % 10 * pow(10, len-i-1);
        x = x / 10;
        i++;
    }
    
    return res;
   
    
    
}