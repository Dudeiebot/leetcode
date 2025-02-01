package heap

import "container/heap"

// keep tracks of all tweets, alll follows and requestid is used to update our timestamp of the tweet
type Twitter struct {
	tweets    map[int]*Tweet
	follows   map[int]map[int]bool
	requestId int
}

// our tweetid and something like a linked list for our previous tweet and here is the timestamp
type Tweet struct {
	tweetId   int
	previous  *Tweet
	timestamp int
}

// first intialization of out Twitter struct
func Constructor() Twitter {
	return Twitter{
		tweets:    make(map[int]*Tweet),
		follows:   make(map[int]map[int]bool),
		requestId: 0,
	}
}

// posting of tweet is straightforward, we check for our user is kind of valid, increment our timestamp upon each tweet
// normal Twitter Instance created
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Init(userId)
	this.requestId += 1
	this.tweets[userId] = &Tweet{
		tweetId:   tweetId,
		previous:  this.tweets[userId],
		timestamp: this.requestId,
	}
}

// Get NewsFeed is the hardest and where out implementation of heap is required, so initially here we check for our user
// upon finishing checking, we check for our followers and if we are following them and taking all the 10 tweet check ing the previous also without storing them
func (this *Twitter) GetNewsFeed(userId int) []int {
	this.Init(userId)

	newsFeed, tweetHeap := []int{}, &TweetHeap{}
	heap.Init(tweetHeap)

	for followeeId, isFollowing := range this.follows[userId] {
		if isFollowing {
			this.Init(followeeId)
			heap.Push(tweetHeap, this.tweets[followeeId])
		}
	}

	for len(*tweetHeap) > 0 && len(newsFeed) < 10 {
		nextTweet := heap.Pop(tweetHeap).(*Tweet)
		if nextTweet.timestamp == 0 {
			break
		}

		newsFeed = append(newsFeed, nextTweet.tweetId)
		heap.Push(tweetHeap, nextTweet.previous)
	}
	return newsFeed
}

// normal follow, remember we are using boolean to
func (this *Twitter) Follow(followerId int, followeeId int) {
	this.Init(followerId)
	this.follows[followerId][followeeId] = true
}

// normal unfollow, remember we are using boolean to check
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.Init(followerId)
	this.follows[followerId][followeeId] = false
}

func (t *Twitter) Init(userId int) {
	if _, hasTweet := t.tweets[userId]; hasTweet {
		return
	}
	t.tweets[userId] = &Tweet{}
	t.follows[userId] = make(map[int]bool)
	t.follows[userId][userId] = true
}

type TweetHeap []*Tweet

func (th TweetHeap) Len() int {
	return len(th)
}

func (th TweetHeap) Less(i, j int) bool {
	return th[i].timestamp > th[j].timestamp
}

func (th TweetHeap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}

func (th *TweetHeap) Push(x any) {
	*th = append(*th, x.(*Tweet))
}

func (th *TweetHeap) Pop() any {
	curr := *th
	val := curr[len(curr)-1]
	*th = curr[:len(curr)-1]

	return val
}
