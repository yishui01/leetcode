/** 
 * Return an array of size *returnSize. 
 * Note: The returned array must be malloced, assume caller calls free(). 
 */  
 
/*
这题目我做不出来，把大神代码复制过来之后看了半天才看懂，递归能用成这样我是真的佩服，后来才知道这TM是一个叫做DFS（深度搜索）的东西，好吧我等下去学习下

主要两点，1、左括号数和右括号数的总数是相等的
         2、只有剩余左括号总数小于等于右括号总数才是有效的序列
         
思路：先拼左括号，拼一个就执行一次递归，递归完成的目标是左括号和右括号都拼完，返回一个有效序列到结果数组中
**/
void generate(char** result, int* Size, int l, int r, char* tmp, int index) {
      printf("此时l是%d--r是%d\n", l, r);
    if(l==0 && r==0) {  
        tmp[index]=0;  
        result[*Size]=(char*)malloc(sizeof(char)*index);  
        strcpy(result[*Size],tmp);  
        (*Size)++;  
        return;  
    }  
    if (l>0) {
        tmp[index]='(';  
        generate(result, Size, l-1, r, tmp, index+1);  
    }  
    if(r>0 && l<r) {  
        tmp[index]=')';  
        generate(result, Size, l, r-1, tmp, index+1);  
    }  
}  
char** generateParenthesis(int n, int* returnSize) {  
    char** result=(char**)malloc(sizeof(char)*100000);  
    char* tmp=(char*)malloc(sizeof(char)*(n*2+1));    
    int l=n,r=n;   
      
    *returnSize=0;    
    generate(result,returnSize,l,r,tmp,0);    
    return result;    
}  