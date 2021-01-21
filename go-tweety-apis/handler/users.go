package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go-tweety/helper"
	"go-tweety/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	//decoding the request
	var req model.SignUp
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err, "decode err")
	}
	fmt.Println(r.Body)
	//encrypting the data
	password := helper.Encrypt([]byte(req.Password), "password")
	consumerKey := helper.Encrypt([]byte(req.ConsumerKey), "apiKey")
	consumerSecret := helper.Encrypt([]byte(req.ConsumerSecret), "apiSecret")
	accessToken := helper.Encrypt([]byte(req.AccessToken), "accessToken")
	accessSecret := helper.Encrypt([]byte(req.AccessSecret), "accessSecret")

	//Load Env file
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGO_URI")

	// Database Config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoURI,
	))
	if err != nil {
		log.Fatal(err)
	}

	client.Connect(ctx)
	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database(os.Getenv("DB_NAME"))
	usersCollection := quickstartDatabase.Collection(os.Getenv("DOCUMENT_NAME"))
	// dateUpdated := time.Now()
	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"username", req.Username}}
	update := bson.D{
		{"$set", bson.D{{"username", req.Username}}},
		{"$set", bson.D{{"password", password}}},
		{"$set", bson.D{{"consumerKey", consumerKey}}},
		{"$set", bson.D{{"consumerSecret", consumerSecret}}},
		{"$set", bson.D{{"accessToken", accessToken}}},
		{"$set", bson.D{{"accessSecret", accessSecret}}}}
	var updatedDocument bson.M
	err = usersCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}

	res := model.SignUpRes{
		Message: req.Username + "Succesfully resigtered.",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

// Displays recent tweets
func ShowTweets(w http.ResponseWriter, r *http.Request) {
	//decoding the request
	var req model.SignUp
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err, "decode err")
	}
	userDetails := helper.GetUserDetails(req.Username, req.Password)
	if !helper.ValidatePassword(req.Password, userDetails.Password) {
		fmt.Println("Incorrect username or password")
	}
	consumerKey := helper.Decrypt(userDetails.ConsumerKey, "apiKey")
	consumerSecret := helper.Decrypt(userDetails.ConsumerSecret, "apiSecret")
	accessToken := helper.Decrypt(userDetails.AccessToken, "accessToken")
	accessSecret := helper.Decrypt(userDetails.AccessSecret, "accessSecret")

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
		len := len(tweets)
		res := make([]model.ShowTweets, len)
		for index, tweet := range tweets {
			res[index].TweetID = tweet.ID
			res[index].Tweet = tweet.Text
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
