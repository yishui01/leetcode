/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
char** letterCombinations(char* digits, int* returnSize) {
    char * curr, ** res;
    int len = strlen(digits), i = 0, j = 1, z = 0, top = 0;
    int level = 0;
    res = (char **) malloc(sizeof(char*)*len*400); //这里有缺陷，可以优化分配内存大小
    char *reflex[] = {
         "",
         "",
        "abc", 
        "def", 
        "ghi", 
        "jkl",
        "mno",
        "pqrs",
        "tuv",
        "wxyz"
    };
    if (len == 0) {
        *returnSize = 0;
        return res;
    } 
    
    int firstlen = strlen(reflex[digits[0]-'0']);
    for(i = 0; i < firstlen; i++) {
        curr = (char*)malloc(sizeof(char)*len+1);
       
        curr[0] = reflex[digits[0]-'0'][i];
        recursion(curr, res, digits, reflex, len, level+1,&top);
    }
    *returnSize = top;
  
    return res;
}

void recursion(char *curr, char ** res, char * digits, char **reflex, int len, int level, int *top) {
    if(level == len){
        int currlen = strlen(curr);
        curr[level] = '\0';
        res[*top] = curr;
         (*top)++;
    }else{
     
           int currlevellen = strlen(reflex[digits[level]-'0']);
            int i = 0;
            for(i = 0; i < currlevellen; i++) {
                char *copy_curr = (char*)malloc(sizeof(char)*len+1);
                strcpy(copy_curr,curr);  
                int currlen = strlen(curr);
                copy_curr[currlen] = reflex[digits[level]-'0'][i];
                recursion(copy_curr, res, digits, reflex, len, level+1, top);
            }
    }
    

    
}
