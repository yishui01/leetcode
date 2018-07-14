int firstMissingPositive(int* nums, int numsSize) {
    //先排序，再遍历分各种情况
    int i = 0, j = 0, z = 0;
    for(i = 0; i < numsSize - 1; i++) {
        for(j = 0; j <numsSize-i-1; j++) {
            if(nums[j] > nums[j+1]) {
                z = nums[j];
                nums[j] = nums[j+1];
                nums[j+1] = z;
            }
        }
    }
    
    for( i = 0; i < numsSize; i++) {
        if(nums[i] <= 0) {
            if(i < numsSize-1) {
                //没到最后一位,看下一位
                if(nums[i+1] <=0){
                    continue;
                } else if(nums[i+1] == 1) {
                    
                    continue;
                } else {
                    return 1;
                }
            } else {
                return 1;
            }
        } else {
            if(i == 0) {
                if(nums[i] > 1) {
                    return 1;
                }
            } else {
                if(nums[i] - nums[i-1] > 1) {
                    if(nums[i-1] <= 0){
                        if(nums[i] > 1) {
                            return 1;
                        } else {
                            continue;
                        }
                    } else {
                        return nums[i-1]+1;
                    }
                }
            }
        }
    }

   return nums[numsSize-1]+1;
   
}