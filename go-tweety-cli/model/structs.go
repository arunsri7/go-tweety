package model

type User struct {
	ID             string `json:_id`
	Username       string `json:username`
	Password       []byte `json:password`
	ConsumerKey    []byte `json:consumerKey`
	ConsumerSecret []byte `json:consumerSecret`
	AccessToken    []byte `json:accessToken`
	AccessSecret   []byte `json:accessSecret`
}
