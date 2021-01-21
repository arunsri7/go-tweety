package cmd

import (
	"fmt"
	"go-tweety/helper"
	"go-tweety/model"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// newTweetCmd represents the newTweet command
var newTweetCmd = &cobra.Command{
	Use:   "newTweet <username> <password>",
	Short: "Sends a tweet given your tweety username and password",
	Run: func(cmd *cobra.Command, args []string) {
		var userDetails model.User
		userDetails = helper.GetUserDetails(args[0], args[1])
		if helper.ValidatePassword(args[1], userDetails.Password) {
			sendTweet(userDetails.ConsumerKey, userDetails.ConsumerSecret,
				userDetails.AccessToken, userDetails.AccessSecret)
		} else {
			fmt.Println("Incorrect username or password")
		}
	},
}

func init() {
	rootCmd.AddCommand(newTweetCmd)
}

// prompts for user input and posts it on twitter
func sendTweet(consumerKey []byte, consumerSecret []byte, accessToken []byte, accessSecret []byte) {
	prompt := promptui.Prompt{
		Label: "Type your tweet",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
	consumerKey = helper.Decrypt(consumerKey, "apiKey")
	consumerSecret = helper.Decrypt(consumerSecret, "apiSecret")
	accessToken = helper.Decrypt(accessToken, "accessToken")
	accessSecret = helper.Decrypt(accessSecret, "accessSecret")

	config := oauth1.NewConfig(string(consumerKey), string(consumerSecret))
	token := oauth1.NewToken(string(accessToken), string(accessSecret))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	_, _, err = client.Statuses.Update(result, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Tweet sent")
	}
}
