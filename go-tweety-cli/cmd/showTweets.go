package cmd

import (
	"fmt"
	"go-tweety/helper"
	"go-tweety/model"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/cobra"
)

// showTweetsCmd represents the showTweets command
var showTweetsCmd = &cobra.Command{
	Use:   "showTweets <username> <password>",
	Short: "Displays your last 20 tweets",
	Run: func(cmd *cobra.Command, args []string) {
		var userDetails model.User
		userDetails = helper.GetUserDetails(args[0], args[1])
		if helper.ValidatePassword(args[1], userDetails.Password) {
			showTweets(userDetails.ConsumerKey, userDetails.ConsumerSecret,
				userDetails.AccessToken, userDetails.AccessSecret)
		} else {
			fmt.Println("Incorrect username or password")
		}
	},
}

func init() {
	rootCmd.AddCommand(showTweetsCmd)
}

// Displays recent tweets
func showTweets(consumerKey []byte, consumerSecret []byte, accessToken []byte, accessSecret []byte) {
	consumerKey = helper.Decrypt(consumerKey, "apiKey")
	consumerSecret = helper.Decrypt(consumerSecret, "apiSecret")
	accessToken = helper.Decrypt(accessToken, "accessToken")
	accessSecret = helper.Decrypt(accessSecret, "accessSecret")

	config := oauth1.NewConfig(string(consumerKey), string(consumerSecret))
	token := oauth1.NewToken(string(accessToken), string(accessSecret))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	// Home Timeline
	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20, //Number of tweets to return
	})

	if err != nil {
		fmt.Println(err)
	} else {
		for index, tweet := range tweets {
			fmt.Println("Tweet no:", index, "--->", tweet.Text, "------- Tweet ID:", tweet.ID)
		}
	}
}
