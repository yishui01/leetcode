char* intToRoman(int num) {
    /*把这整数从个位一位一位的截取,转换成字符串后从尾部往前插入数组*/
    int len = 0, remain = 0, bak = 0, i = 0, full = 20, index = 18;//长度、余数、新数字
    char curr, five, ten;
    char * dst = (char *)malloc(sizeof(char)*full);
    while(num > 0) {
         remain = num % 10;
         if(len == 0){
                curr = 'I';
                five = 'V';
                ten = 'X';
            } else if(len == 1) {
                curr = 'X';
               five = 'L';
                ten = 'C';
            } else if(len == 2) {
                curr = 'C';
               five = 'D';
               ten = 'M';
            } else if(len == 3) {
                curr = 'M';
             
            }
        if(remain >= 1 && remain <=3){
                for(i = 0; i<remain; i++) {
                    dst[index] = curr;
                    index--;
                }
        } else if(remain == 4) {
                    dst[index] = five;
                    index--;
                    dst[index] = curr;
                    index--;
        } else if(remain >=5 && remain <=8) {
             
            for(i = 0; i<remain-5; i++) {
                    dst[index] = curr;
                    index--;
                }
                    dst[index] = five;
                    index--;
        }else if(remain == 9) {
                    dst[index] = ten;
                    index--;
                    dst[index] = curr;
                    index--;
        }else{
            //这里余数是0的时候了，不用管
        }
        num /= 10;
        len++;
    }
    index++;
    dst[full-1] = '\0';
    return dst+index;
}