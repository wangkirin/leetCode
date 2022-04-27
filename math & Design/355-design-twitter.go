package main

import (
	"sort"
)

func main() {

}

type Twitter struct {
	seq         int
	usertweets  map[int][][2]int
	userfollows map[int][]int
}

func Constructor() Twitter {
	return Twitter{
		seq:         0,
		usertweets:  make(map[int][][2]int),
		userfollows: make(map[int][]int),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.usertweets[userId] = append(this.usertweets[userId], [2]int{this.seq, tweetId})
	this.seq++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	follows := append(this.userfollows[userId], userId)
	tweets := [][2]int{}
	for _, follow := range follows {
		tweets = append(tweets, this.usertweets[follow]...)
	}
	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i][0] > tweets[j][0]
	})
	n := 10
	if len(tweets) < 10 {
		n = len(tweets)
	}
	ans := []int{}
	for i := 0; i < n; i++ {
		ans = append(ans, tweets[i][1])
	}
	return ans
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	for _, id := range this.userfollows[followerId] {
		if id == followeeId {
			return
		}
	}
	this.userfollows[followerId] = append(this.userfollows[followerId], followeeId)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	sort.Ints(this.userfollows[followerId])
	index := sort.SearchInts(this.userfollows[followerId], followeeId)
	if len(this.userfollows[followerId]) == 1 {
		this.userfollows[followerId] = nil
	} else if len(this.userfollows[followerId]) == 0 {
		return
	} else {
		this.userfollows[followerId] = append(this.userfollows[followerId][:index], this.userfollows[followerId][index+1:]...)
	}
}
