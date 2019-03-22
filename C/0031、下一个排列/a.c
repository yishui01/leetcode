void nextPermutation(int* nums, int numsSize) {
    int i = 0, j = 0, z = 0,min = 0,min_index = 0, tmp = 0, ismax = 1;
    //看看是否是最大排序
    for(i = 0; i < numsSize-1; i++) {
        if (nums[i] < nums[i+1]) {
           ismax = 0;
            break;
        }
    }
    if (ismax) {
       //如果是最大排序，翻转数组返回
        for(i = 0; i < numsSize/2; i++) {
            tmp = nums[i];
            nums[i] = nums[numsSize-i-1];
            nums[numsSize - i-1] = tmp;
        }
    } else {
        //从后往前每两个数比较一次，如果后面的数大于前面的数，准备交换，如果后面的数的后面还有比它小的数，那个最小的数如果比交换为的数字大
        //就是真正要替换的数
        //例如1,3,3,2  发现3比1大，但是3后面还有两个数分别为3,2 而2比准备交换的3这个数字还小，且2是后面这些数字中最小的且比交换位的数字1大，
        //所以真正要替换的是2这个数字
        //将2和1替换之后，变成了2,3,3,1 这是不对的，所以要将2后面的数字按照从小到大排序
        //最终结果变成了2,1,3,3     it's good HAHAHA
        for(i = numsSize-1; i > 0; i--) {
            if(nums[i] > nums[i-1]) {
                printf("到了");
                //后面的数字比前面的大
                min = nums[i];
                min_index = i;
                for(j = i; j < numsSize ; j++) {
                    //找出后半段最小的那个数字且要比即将换位的数字要大，比如2,3,1
                    if(nums[j] < min && nums[j] > nums[i-1]) {
                        min = nums[j];
                        min_index = j;
                    }
                }
                
                //此时nums[j]才是真正要交换的值
                tmp = nums[i-1];
                nums[i-1] = nums[min_index];
                nums[min_index] = tmp;
                //再把后半段按照从小到大排序
                for(j = i; j < numsSize; j++) { 
                    for(z = i; z < numsSize - 1; z++) {
                        //注意这里的冒泡排序，外层初始值不为0，需要注意内层循环条件不能为z < numsSize - j - 1 ，j初始值不为0，不能减j
                        if(nums[z] > nums[z+1]) {
                            tmp = nums[z];
                            nums[z] = nums[z+1];
                            nums[z+1] = tmp;
                        }
                    }
                }
                break;
            }
        }
            
        
    }
    
}