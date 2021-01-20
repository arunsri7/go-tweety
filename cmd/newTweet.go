package cmd

import (
	"context"
	"fmt"
	"go-tweety/helper"
	"go-tweety/model"
	"log"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// newTweetCmd represents the newTweet command
var newTweetCmd = &cobra.Command{
	Use:   "newTweet <username> <password>",
	Short: "Sends a tweet given your tweety username and password",
	Run: func(cmd *cobra.Command, args []string) {
		getUserDetails(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(newTweetCmd)
}

//getUserDetails returns the object from the db given the username and password
func getUserDetails(username string, password string) {
	// Database Config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Load Env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoURI,
	))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	//Cancel context to avoid memory leak

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	defer client.Disconnect(ctx)
	quickstartDatabase := client.Database("goTweety")
	productsCollection := quickstartDatabase.Collection("users")
	var userDetails model.User
	if err = productsCollection.FindOne(ctx,
		bson.M{"username": username}).Decode(&userDetails); err != nil {
		fmt.Println("wrong username or password")
	}
	if validatePassword(password, userDetails.Password) {
		sendTweet(userDetails.ConsumerKey, userDetails.ConsumerSecret,
			userDetails.AccessToken, userDetails.AccessSecret)
	} else {
		fmt.Println("Incorrect Username or Password")
	}
}

//decrypt and validate user password
func validatePassword(password string, mongoDBPass []byte) bool {
	decryptedPass := helper.Decrypt(mongoDBPass, "password")
	if password == string(decryptedPass) {
		return true
	}
	return false

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
