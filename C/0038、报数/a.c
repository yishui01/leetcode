char* countAndSay(int n) {
    if(n == 1) {
        return "1";
    } else {
        char *p;
        p = countAndSay(n-1); //接收传递过来的值
        int len = strlen(p), now_count = 0, now_index = 0;
        char now_char = '\0';
        char * tmp = (char *)malloc(sizeof(char) * len * 2);
        while(*p) {
            now_char = *p;
            now_count = 0;
            while(*p == now_char) {
                now_count++;
                p++;
            }
            tmp[now_index++] = now_count + '0';
            tmp[now_index++] = now_char;
        }
        tmp[now_index] = '\0';
        return tmp;
    }

}