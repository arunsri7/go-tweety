package cmd

import (
	"fmt"
	"go-tweety/helper"
	"go-tweety/model"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/cobra"
)

// deleteTweetsCmd represents the deleteTweets command
var deleteTweetsCmd = &cobra.Command{
	Use:   "deleteTweets <username> <password> <tweetID>",
	Short: "Deletes the tweet with the given tweet id. Tweet Id can be obtained from show tweets command",
	Run: func(cmd *cobra.Command, args []string) {
		var userDetails model.User
		userDetails = helper.GetUserDetails(args[0], args[1])
		if helper.ValidatePassword(args[1], userDetails.Password) {
			tweetID, _ := strconv.ParseInt(args[2], 10, 64)
			deleteTweets(tweetID, userDetails.ConsumerKey, userDetails.ConsumerSecret,
				userDetails.AccessToken, userDetails.AccessSecret)
		} else {
			fmt.Println("Incorrect username or password")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteTweetsCmd)
}

// deletes the tweet given the tweet id
func deleteTweets(tweetID int64, consumerKey []byte, consumerSecret []byte, accessToken []byte, accessSecret []byte) {
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
	tweet, _, err := client.Statuses.Destroy(tweetID, &twitter.StatusDestroyParams{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("deleted the tweet with ID %d sucessfully", tweet.ID)
	}
}
