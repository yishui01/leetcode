package main

//参考大神解法，O（n）一遍过，你敢信？
func lengthOfLongestSubstring(s string) int {
	//一排格子，下雨，雨点滴到格子里，看格子里之前有没有水，有的话就更新水的标志为当前的i，每次滴雨都计算当前的max
	//计算max的方式就是当前的i减去起点的i，起点的i每次在滴到重复的格子时会被更新成那个格子里的水的i。初始起点（start）为
	//-1，-1的好处就是当i为0时，以上max计算公式任然成立，当s长度为1时，i-start，max = 0-（-1） = 1,膜拜大神
	tmp := make([]int, 256)
	max := 0
	start := -1
	for k, _ := range tmp {
		tmp[k] = -1
	}
	for i := 0; i < len(s); i++ {
		if tmp[int(s[i])] > start {
			//找到重复的了
			start = tmp[int(s[i])]
		}
		tmp[int(s[i])] = i //赋值占位，然后原来有值的也会被赋值为新的值
		if max < i-start {
			max = i - start
		}
	}

	return max
}
/***********************************下面是本人的菜鸟解法，是的，菜到极致 ╮(╯▽╰)╭ **********************************/
//func lengthOfLongestSubstring(s string) int {
//	if len(s) <= 1 {
//		return len(s)
//	}
//
//	tmp := make([]int, len(s))
//	max := 0
//	now_index := 0
//	for _, v := range s {
//		if find_index := isExists(tmp, int(v)); find_index != -1 {
//			//存在就把当前的数的后面那个数放到头部，开始计数 dvdf 遇到第二个d时，把第一个d后面的那个数放到头顶
//			if now_index > max {
//				max = now_index
//			}
//			// fmt.Println(now_index,"--",find_index)
//			// fmt.Println(tmp)
//			tmp[now_index] = int(v)
//			now_index = now_index - find_index
//
//			for i := 0; i < len(tmp); i++ {
//				if find_index+1+i >= len(tmp) {
//					tmp[i] = 0
//				} else {
//					tmp[i] = tmp[find_index+1+i]
//				}
//			}
//			tmp[now_index] = int(v)
//			//  fmt.Println(tmp)
//
//		} else {
//			//不存在就把字符追加到数组后面
//			tmp[now_index] = int(v)
//			now_index++
//		}
//
//	}
//	if now_index > max {
//		max = now_index
//	}
//	return max
//}
//
//func isExists(sli []int, num int) int {
//	res := -1
//	for k, v := range sli {
//		if v == num {
//			res = k
//			return res
//		}
//	}
//
//	return res
//}
