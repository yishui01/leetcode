/****

直接看解释，很容易理解，原理：利用快排，只要找到一个主元，他的右边有k-1个元素的话，那么这个点是不是就是第k大的元素？

Let's say we have an array nums as below and k = 2:
[4, 7, 5, 9, 2, 3]

First we chose a pivot. For simplicity, I choose nums[0] = 4. You can also choose nums[end] or nums[mid], if you want.
After choosing the pivot, we organize the array so that the array looks like [(less than 4), 4, (greater than 4)].
If we perform a normal quick sort, then we will recursively run the same code for the former half and the latter half of nums.

But for solving this problem, we should stop here and think what is happening now.
We can't guarantee the left side and the right side is sorted, but we can guarantee all numbers left side of 4 is less than 4 and all ones right side of 4 is greater than 4.
Then, if there is only 1 number at the right side of 4, can't we say 4 is k = 2 nd largest number in nums?

And even if 4 is not k th largest number in nums, like binary search, we need to only either side of 4 for next quick sort. We aren't interested in sorting, we just are interested in which number is k th largest.
So, before finding k th element, we need to only swap k times elements. That is to say, we can achieve expected O(k) complexity, while in worst case it's still O(n^2).
***/


func findKthLargest(nums []int, k int) int {
    return doFindKthLargest(nums, k, 0, len(nums)-1)
}

func doFindKthLargest(nums []int, k int, start, end int) int {
    nlen := len(nums)
    targetPos := nlen - k //第k大的元素在有序数组中的下标为 nlen - k
    
    for {
        pivotIndex := partition(nums, start, end)
        if pivotIndex == targetPos {
            return nums[pivotIndex]
        } else if pivotIndex < targetPos { //如果当前主元的位置比目标位置小，那么很明显，target在右半段
            start = pivotIndex + 1
        } else { //反之在左半段
            end = pivotIndex - 1
        }
    }
}






//单路快排，可以优化成双指针去重复的快排
// func partition(nums []int, left, right int) int {
//     //以头部元素为标志进行区分
// 	j := left
// 	for i := left + 1; i <= right; i++ {
// 		if nums[i] <= nums[left] {
// 			j++
// 			nums[j], nums[i] = nums[i], nums[j]
// 		}
// 	}
// 	//最后把标志位放到对应的分割点位置,这里是从小到大排序，将left和j号位的元素互换,现在j指向的是前半段最后一个元素
// 	nums[left], nums[j] = nums[j], nums[left]
// 	return j
// }

func partition(arr []int, begin, end int) int {
	pivot := arr[begin]
	for begin < end {
		/*
		 * 从后往前遍历，找到比pivot小的值
		 */
		for begin < end && arr[end] >= pivot {
			end--;
		}
		arr[begin] = arr[end];
		/*
		 * 从前往后遍历，找到比pivot大的值；
		 */
		for begin < end && arr[begin] <= pivot {
			begin++
		}
		arr[end] = arr[begin];
	}
	arr[begin] = pivot;
	return begin;
}
