/*左括号入栈，右括号就往前面找左括号，找到了就消消乐一样将左右括号替换为' '，最后找出数组中最大的连续空字符长度就是结果*/
int longestValidParentheses(char* s) {
    int max_len = 0, i = 0, j = 0, tmp_len = 0, slen = strlen(s);
    char *tmp = (char *)malloc(sizeof(char) * slen);
    memset(tmp, 'A', slen);
    while(*s) {
        if(*s == '(') {
            //左括号入栈
           tmp[i] = '('; 
        } else {
            //右括号则找出最近的左括号，清空
            for(j = i-1; j  >= 0; j--) {
                if(tmp[j] == ' '){
                    continue;
                } else if(tmp[j] == '(') {
                    //找到左括号了，清空左右括号
                    tmp[j] = ' '; 
                    tmp[i] = ' ';
                    break;
                } else {
                    //找到了个右括号
                    break;
                }
            }
        }
        i++;
        s++;
    }
    //找出数组中最大长度的连续空字符
    for(i = 0; i < slen; i++) {
        if(tmp[i] == ' ') {
            tmp_len++;
        } else {
            if(tmp_len > max_len) {
                max_len = tmp_len;
            }
            tmp_len = 0;
        }
    }
    free(tmp);
    if(tmp_len > max_len) {
        max_len = tmp_len;
    }
    return max_len;
}