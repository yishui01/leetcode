/*Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].*/

#include <stdio.h>
#include <stdlib.h>
int* twoSum(int* nums, int numsSize, int target);

int main()
{
	int test[4] = { 2,7,11,15};
	int * res = twoSum(test,4, 18 );
	printf("%d---%d", res[0], res[1]);
	getchar();
	return 0;
}

int* twoSum(int* nums, int numsSize, int target) {
	//两个数组中每两个数字相加一次即可
	int i, j;
        int *return_arr = (int *) malloc(2 * sizeof(int));
	for (i = 0; i<numsSize; i++) {
		for (j = i + 1; j< numsSize; j++) {
			if (nums[i] + nums[j] == target) {
				return_arr[0] = i;
				return_arr[1] = j;
				return return_arr;
			}
		}
	}

	return return_arr;
}
