 
//传说中的深度优先搜索，学到了

//交换两个指针的值
void swap_num(int *a, int *b)
 {
     int temp = *a;
     *a = *b;
     *b = temp;
 }
 

 void permut(int **result, int *nums, int end, int *rows, int beg)
 {
     int i;
     if (beg == end) //已经是最后一个元素，递归返回
     {
         //从最后一层开始产出元素到结果数组中，逐级回归
         (*rows)--;
         for(i = 0; i != end; i++)
            result[(*rows)][i] = nums[i];
     }
     else
     {
         for (i = beg; i < end; i++)
         {
             swap_num(&nums[beg], &nums[i]);
             permut(result,nums, end, rows, beg+1);
             swap_num(&nums[beg], &nums[i]);
         }
     }
     
 }
 
int** permute(int* nums, int numsSize, int* returnSize) {
    int **result, i;
    *returnSize = 1;
    /* 计算全排列数 */
    for(i = numsSize; i > 1; i--)
    {
        (*returnSize) *= i;
    }
    /* 为行分配内存 */
    result = (int **)malloc((*returnSize)*sizeof(int *)); //6 行
    /* 为列分配内存 */
    for (i = 0; i != (*returnSize); i++)
        *(result + i) = (int *)malloc(numsSize * sizeof(int)); //3列
        
        int rows = *returnSize;
        permut(result, nums, numsSize, &rows, 0);
        return result;
    
}