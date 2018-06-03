int romanToInt(char* s) {
    int res = 0;
    while(*s) {
        if(*s == 'I') {
            res++;
        } else if(*s == 'X') {
            res+=10;
            if(*(s-1) == 'I') {
                res -= 2;
            }
        } else if(*s == 'C') {
            res+=100;
            if(*(s-1) == 'X') {
                res -= 20;
            }
        } else if(*s == 'M') {
            res+=1000;
            if(*(s-1) == 'C') {
                res -= 200;
            }
        } else if(*s == 'V') {
            res+=5;
            if(*(s-1) == 'I') {
                res -= 2;
            }
        } else if(*s == 'L') {
            res +=50;
            if(*(s-1) == 'X') {
                res -= 20;
            }
        } else if(*s == 'D') {
            res += 500;
            if(*(s-1) == 'C') {
                res -= 200;
            }
        }
        s++;
    }
    
    return res;
}