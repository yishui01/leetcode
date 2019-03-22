bool isValid(char* s) {
     //构造一个栈结构数组即可
    if(*s == '\0')return true;
    char left[10000];
    int left_i = 0;
    
    while(*s) {
        if(*s == '(' || *s == '[' || *s == '{') {
            //左括号入栈
            left[left_i] = *s;
            left_i++;
        } else if(*s == ')' || *s == ']' || *s == '}') {
            //右括号出栈
            if(left_i == 0){
                return false;
            } else {
               switch(*s){
                    case ')':
                        if(left[left_i-1] != '('){
                            return false;
                        }
                        break;
                    case ']':
                        if(left[left_i-1] != '['){
                            return false;
                        }
                        break;
                    case '}':
                        if(left[left_i-1] != '{'){
                            return false;
                        }
                        break;
                    case '\0':
                        continue;
                        break;
                }
                left_i--;
               
            }
        } else if(*s == '\0') {
            continue;
        } else {
            return false;
        }
       
        s++;
    }
    //左边的没弹完直接GG
    if(left_i!=0)return false;
    
    return true;
}