char* multiply(char* num1, char* num2) {
    if(((*num1)-'0') == 0 || ((*num2)-'0') == 0)return "0";
    int len1 = strlen(num1), len2 = strlen(num2), i = 0, j = 0, size = 300, isfull = 0;
    char *res = (char *)malloc(sizeof(char)*size);
    memset(res, 0, size);
    res[size-1] = '\0';
    for(i = 0; i < len1; i++) {
        for(j = 0; j < len2; j++) {
            res[size-2-i-j] += (*(num1+len1-1-i) - '0') * (*(num2+len2-1-j) - '0');
            res[size-2-i-j] += isfull; //加上之前的进位
            isfull = res[size-2-i-j] / 10; //计算新的进位值
            res[size-2-i-j] %=10; 
        }
        if(isfull > 0) {
            //最后一轮还有进位的话，补上去
            res[size-2-i-j] = isfull;
        }
        isfull = 0;
    }
    
    
    //找到开始的那个点，再把所有的数字转换为字符
    i = 0;
    while(*(res+i) == 0) {
        i++;
    }

    //现在res+i就是起点
    j = 0;
    while(((res+i)+j)<res+size-1) { //两个指针相比较，把数字转换为字符
        *((res+i)+j) += '0';
        j++;
    }
    
    return res+i;
    
}