package main

import (
	"go-tweety/cmd"
)

func main() {
	cmd.Execute()
	// consumerKey := "866BamNTPZbBy0nNN1bE11MJE"
	// consumerSecret := "nHhYNZN5Fmexa7REBGsN1I3WzCOWj6X9y7ofNqE6BUOHYLtAN4"
	// accessToken := "1304015396710658048-LvE0kCyu8eTJLOxBIUHO2RjCYy4T4t"
	// accessSecret := "l2BohcWYGx5ATA9PppePmZfDC8fp7yHWUXibES5q57OjL"

	// config := oauth1.NewConfig(consumerKey, consumerSecret)
	// token := oauth1.NewToken(accessToken, accessSecret)
	// httpClient := config.Client(oauth1.NoContext, token)

	// // Twitter client
	// client := twitter.NewClient(httpClient)

	// // Home Timeline
	// _, resp, _ := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
	// 	Count: 20,
	// })
	// fmt.Println(resp)

	// // Send a Tweet
	// _, response, _ := client.Statuses.Update("just setting up my twitter", nil)

	// fmt.Println(response)
	// Status Show
	// tweet, resp, err = client.Statuses.Show(585613041028431872, nil)

	// // Search Tweets
	// search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
	// 	Query: "gopher",
	// })

	// // User Show
	// user, resp, err := client.Users.Show(&twitter.UserShowParams{
	// 	ScreenName: "dghubble",
	// })

	// // Followers
	// followers, resp, err := client.Followers.List(&twitter.FollowerListParams{})
}
