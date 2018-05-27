int reverse(int x) {
    //先确定整数有几位
    int len = 1, res = 0, copy;
    copy = x;
    if (x <0){
        if(-x < 0) {
            return 0;
        }
        copy = -x;
    }
 
    while((copy / 10) > 0) {
        len ++;
        copy = copy /10;
        
    }
   
    if(len == 1) {
        return x;
    }
    
    int i = 0, curren = 0, pre = 0, s;
    for (i = 0; i < len; i++) {
        s = x / pow(10, len-1-i);
        curren = s-pre;
        pre = s*10;
        if(x > 0) {
            if (pow(2,31)-1-res < curren*pow(10,i)){
                return 0;
            }
        } else {
            if (pow(-2,31)-1-res > curren*pow(10,i)){
                return 0;
            }
        }
        res += curren*pow(10,i);
    }

    return res;
    
    
}