void recurFind(int* candidates, int candidatesSize, int target, int **res, 
                int** columnSizes, int* returnSize, int stack[], int top, int cursum) {

    if (cursum > target) {
        return;
    }
    if (cursum == target) {
        res[(*returnSize)] = malloc(top * sizeof(int));
        (*columnSizes)[(*returnSize)] = top; //结果数组每个元素的大小,这里是二维指针，先取一次值再赋值
        int i;
        for (i = 0 ; i < top ; i++) {
            res[*returnSize][i] = stack[i]; //吧临时数组放到结果数组中
        }
        (*returnSize)++;
        return;
    }
    int i;
    for (i = 0 ; i < candidatesSize ; i++) {
        
        if (top > 0 && candidates[i] < stack[top-1]) continue; //这一步实际上是去重复，默认是把传进来的数组当成从小到大排列，stack中的最后一个数字比i的键值大的话，直接跳过，因为键值对应的那一层最外层循环已经全部搜索完毕，比如输入[2,3,6,7],最开始以2为起点的，搜索完了之后回到最外层，现在是以3为起点，那么组合中就不需要2这个数字了，因为所有有关于2的组合已经在之前全部找出
        stack[top++] = candidates[i];
        recurFind(candidates, candidatesSize, target, res, columnSizes, returnSize, stack, top, cursum + candidates[i]);
        top--;
    }
}
int** combinationSum(int* candidates, int candidatesSize, int target, int** columnSizes, int* returnSize) {
    int stack[1000];
    int ** res = malloc(1000 * sizeof(int *));
    * columnSizes = malloc(1000 * sizeof(int)); 
    * returnSize = 0;
    int cursum = 0, top = 0;
    recurFind(candidates, candidatesSize, target, res, columnSizes, returnSize, stack, top, cursum);
    return res;
}