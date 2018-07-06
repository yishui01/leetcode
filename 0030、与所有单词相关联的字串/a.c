/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* findSubstring(char* s, char** words, int wordsSize, int* returnSize) {
    //先遍历看下子串长度是否相同，不同直接GG
    int i = 0, j = 0, z = 0, rescount = 0, flag = 0, is_bad = 0, word_len = 0, s_len = 0, all_len = 0; 
    //word_len为每个单词长度、s_len为查询字符串长度、all_len为所有单词总长度
    int *res = (int *)malloc(sizeof(int)*1000);
    
    for(i = 0; i < wordsSize; i++) {
        if(i == 0) {
            word_len = strlen(words[i]);
        } else {
            if(word_len != strlen(words[i]))return NULL;
        }
    }
    if(!word_len){
        return NULL;
    }
    
    all_len = wordsSize * word_len;
    s_len = strlen(s);
    int have_len = s_len - all_len + 1; //需要遍历次数
    if(have_len < 0)return NULL;
    int *tmp = (int *)malloc(sizeof(int)*wordsSize); //用于记录每个单词出现的次数，用下标对应
    char *tmp_w = (char *)malloc(sizeof(char)*word_len+1); //用于记录临时比对字符串
    tmp_w[word_len] = '\0';
    char *tmp_s;
    //解法：
    //逐个字符遍历s，直到剩余长度小于所有单词总长度。
    //每当进入下一个字符，都进行一个内层循环，内层循环为以当前字符为起点，每次遍历单词的长度，看这个单词是否存在
    //存在时，判断当前这个单词出现了几次，超过一次直接GG
    //不存在时也直接GG
    //总共遍历次数为单词数目，遍历完如果如何条件，则将此时s的字符下标存入结果中
    
    i = 0;
    while(i < have_len) {
         memset(tmp,0,sizeof(int)*wordsSize);
        for(j = 0; j < wordsSize; j++) {
            tmp_s = s+j*word_len;
            flag = 0;
            for(z = 0; z < word_len; z++) {
                tmp_w[z] = tmp_s[z];
            }
            //先比对，看是否存在,存在的话再去看是否出现次数为0
            for(z = 0; z < wordsSize; z++) {
                if (strcmp(tmp_w, words[z]) == 0) {
                    if(tmp[z] == 0) {
                        tmp[z]++;
                        flag = 1;
                        break; //占了一个位置就马上跳出
                    }
                }
                
            }
            
            if(!flag){
                flag = 0;
                break;
            }
        }
        if(flag) {
            for(z = 0; z < wordsSize; z++) {
                if(tmp[z] != 1){
                    is_bad = 1;
                    break;
                }
            }
            if(!is_bad) {
                res[rescount] = i;
                rescount++;
            }
        }
       
        is_bad = 0;
        s++;
        i++;
    }
    *returnSize = rescount;
    
    return res;
    
}