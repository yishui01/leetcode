int romanToInt(char* s) {
    int res = 0;
    while(*s) {
        switch(*s) {
            case 'I':
                res++;
                break;
            case 'X':
                 res+=10;
                if(*(s-1) == 'I') {
                    res -= 2;
                }
                break;
            case 'C':
                 res+=100;
                if(*(s-1) == 'X') {
                    res -= 20;
                }
                break;
            case 'M':
                res+=1000;
                if(*(s-1) == 'C') {
                    res -= 200;
                }
                break;
            case 'V':
                res+=5;
                if(*(s-1) == 'I') {
                    res -= 2;
                }
                break;
            case 'L':
                res +=50;
                if(*(s-1) == 'X') {
                    res -= 20;
                }
                break;
            case 'D':
                res += 500;
                if(*(s-1) == 'C') {
                    res -= 200;
                }
                break;
        }

        s++;
    }
    
    return res;
}
