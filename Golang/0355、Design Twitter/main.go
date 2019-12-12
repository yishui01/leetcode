type Twitter struct {
    follows map[int]map[int]bool //全局列表 userId->自己关注了哪些人
    posts map[int][][2]int //全局文章列表，userId->[postId,Nanosecond]
}


/** Initialize your data structure here. */
func Constructor() Twitter {
    return Twitter{
        follows:make(map[int]map[int]bool),
        posts:make(map[int][][2]int),
    }
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.posts[userId] = append(this.posts[userId],[2]int{tweetId,time.Now().Nanosecond()})
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
    //找出自己和关注的人的ID
    if this.follows[userId] == nil {
        this.follows[userId] = make(map[int]bool)
        this.follows[userId][userId] = true
    }
    res := make([][2]int,0)
    for k,v := range this.posts {
        if _,ok := this.follows[userId][k];ok {
            res = append(res,v...)
        }
    }

    sort.Slice(res,func(i,j int)bool{
        return res[i][1] > res[j][1]
    })

    sortRes := make([]int,len(res))

    for k,v := range res {
        sortRes[k] = v[0]
    }
    if len(sortRes) > 10 {
        sortRes = sortRes[:10]
    }
    return sortRes
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.follows[followerId] == nil {
         this.follows[followerId] = make(map[int]bool)
         this.follows[followerId][followerId] = true
    }
    this.follows[followerId][followeeId] = true
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if followerId == followeeId {
        return 
    }
     if this.follows[followerId] != nil {
           delete(this.follows[followerId],followeeId)
     }
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
