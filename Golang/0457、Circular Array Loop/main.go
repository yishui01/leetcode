func circularArrayLoop(nums []int) bool {
    //1、对每个点dfs，记录走过的轨迹,先标记为*，代表是本次轨迹，本次完成后再将本次轨迹标记为#,本次走到历史轨迹直接返回success，否则把本次轨迹所有的点都放到map中，下次遇到这样的点就直接break
    //2、记录方向，方向变化，方向变了直接break
    //3、看看最后能不能回到本次路径上的某个点
    nowMap := make(map[int]bool) //记录本次轨迹
    history := make([]int,len(nums)) //历史轨迹
    for k,v := range nums {
        oldDirect := (v >= 0)
        oldIndex := k
        nowMap[oldIndex] = true
        for {
            newIndex := oldIndex+nums[oldIndex]%len(nums)
            if newIndex < 0 {
                newIndex = len(nums)+newIndex
            }
            if newIndex >= len(nums) {
                newIndex = newIndex-len(nums)
            }
            if newIndex == oldIndex { //原地踏步，break
                break
            }
            newDirect := (nums[newIndex] >= 0)
            if newDirect != oldDirect { //反向了，break
                break
            }
            if history[newIndex] == 1 { //到了一个不行的点，break
                break
            }
            if nowMap[newIndex]{ //到了本次轨迹上的点 完美循环
                //回到本次路径的某个点了
                return true
            }

            nowMap[newIndex] = true //存入map，方便直接取出
            oldIndex = newIndex
        }

        //不通过,加入history，清空now
        for k,_ := range nowMap {
            history[k] = 1
        }
        nowMap = make(map[int]bool)
    }

    return false
}
