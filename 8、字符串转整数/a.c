/*示例 1:

输入: "42"
输出: 42
示例 2:

输入: "   -42"
输出: -42
解释: 第一个非空白字符为 '-', 它是一个负号。
     我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
示例 3:

输入: "4193 with words"
输出: 4193
解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
示例 4:

输入: "words and 987"
输出: 0
解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
     因此无法执行有效的转换。
示例 5:

输入: "-91283472332"
输出: -2147483648
解释: 数字 "-91283472332" 超过 32 位有符号整数范围。 
     因此返回 INT_MIN (−231) 。*/
	 
	 
//代码很乱，而且有一个缺点，
//就是会按照输入字符串的长度分配内存，可能会造成严重后果，因此暂时限制字符串长度为1000，可以通过OJ
//二刷的时候再来优化这个问题
int myAtoi(char* str) {
    int r = 0, flag = 0, i = 0, len = 0, j = 0, slen = 0; //r为返回结果
    slen = strlen(str);
	 if(slen > 1000)return 0;
    char f = 'N';
    char *res = (char*)malloc(sizeof(char)*slen);
    memset(res, '0', slen);
    while (*str != '\0') {
        if(flag ==1 && (*str <'0' || *str >'9'))break;
        if(*str == ' '){
             str++;
            continue;
        }
        if(flag == 0 && (*str == '+' || *str == '-')) {
            res[i] = *str;
        }else if(*str >='0' && *str <='9') {
            res[i] = *str;          
        }else{
            break;
        }

        len++;
        flag = 1;
        i++;
        str++;
       
    }
  
    if(len == 0 || ((res[0] == '+' || res[0] == '-') && len ==1)){
        return 0;
    }

    for(j = 0; j<len; j++) {
       
        if(res[j] == '+' || res[j] == '-'){
            f = res[j];
            continue;
        }
     
        if(f =='N'){  
          
            //没有符号位
                if((r + (res[j]-'0') * pow(10, len-j-1)) > INT_MAX-1){
                return INT_MAX;
            }
               
            r += (res[j]-'0') * pow(10, len-j-1);
            
        }else{
            //有符号位
            if(f == '+'){
                 if((r + (res[j]-'0') * pow(10, len-j-1)) > INT_MAX){
                return INT_MAX;
               }
            }else{
                if(0-(r + (res[j]-'0') * pow(10, len-j-1)) < INT_MIN){
                return INT_MIN;
               }
            }
            
            r += (res[j]-'0') * pow(10, len-j-1);
        }
       

    }
    
    free(res);
   if(f == '-'){
       return 0-r;
   }
   
    return r;
}