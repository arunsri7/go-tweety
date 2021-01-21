package model

type SignUp struct {
	Username       string `json:username`
	Password       string `json:password`
	ConsumerKey    []byte `json:consumerKey`
	ConsumerSecret []byte `json:consumerSecret`
	AccessToken    []byte `json:accessToken`
	AccessSecret   []byte `json:accessToken`
}

type SignUpRes struct {
	Message string
}

type User struct {
	ID             string `json:_id`
	Username       string `json:username`
	Password       []byte `json:password`
	ConsumerKey    []byte `json:consumerKey`
	ConsumerSecret []byte `json:consumerSecret`
	AccessToken    []byte `json:accessToken`
	AccessSecret   []byte `json:accessSecret`
}

type ShowTweets struct {
	Tweet   string
	TweetID int64
}

type NewTweet struct {
	Tweet    string `json:tweet`
	Username string `json:username`
	Password string `json:password`
}

type DeleteTweet struct {
	TweetID  int64  `json:tweetid`
	Username string `json:username`
	Password string `json:password`
}
