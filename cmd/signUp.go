package cmd

import (
	"context"
	"go-tweety/helper"
	"log"
	"time"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// signUpCmd represents the signUp command
var signUpCmd = &cobra.Command{
	Use:   "signUp",
	Short: "Add your api keys and access token keys to sign up ",
	Run: func(cmd *cobra.Command, args []string) {
		password := helper.Encrypt([]byte(args[1]), "password")
		consumerKey := helper.Encrypt([]byte(args[2]), "apiKey")
		consumerSecret := helper.Encrypt([]byte(args[3]), "apiSecret")
		accessToken := helper.Encrypt([]byte(args[4]), "accessToken")
		accessSecret := helper.Encrypt([]byte(args[5]), "accessSecret")
		// fmt.Println(password, consumerKey, consumerSecret, accessToken, accessSecret)
		signUp(args[0], password, consumerKey, consumerSecret, accessToken, accessSecret)
		// fmt.Println("Starting the application...")
		// ciphertext := helper.Encrypt([]byte("Hello World"), "password")
		// fmt.Printf("Encrypted: %x\n", password)
		// plaintext := helper.Decrypt(ciphertext, "password")
		// fmt.Printf("Decrypted: %s\n", plaintext)
		// consumerKey := "866BamNTPZbBy0nNN1bE11MJE"
		// consumerSecret := "nHhYNZN5Fmexa7REBGsN1I3WzCOWj6X9y7ofNqE6BUOHYLtAN4"
		// accessToken := "1304015396710658048-LvE0kCyu8eTJLOxBIUHO2RjCYy4T4t"
		// accessSecret := "l2BohcWYGx5ATA9PppePmZfDC8fp7yHWUXibES5q57OjL"
	},
}

func init() {
	rootCmd.AddCommand(signUpCmd)
}

func signUp(username string, password []byte, consumerKey []byte, consumerSecret []byte,
	accessToken []byte, accessSecret []byte) {
	// Database Config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		`mongodb+srv://admin:admin@cluster0.1vpl5.mongodb.net/goTweety?retryWrites=true&w=majority`,
	))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("goTweety")
	usersCollection := quickstartDatabase.Collection("users")
	// dateUpdated := time.Now()
	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"username", username}}
	update := bson.D{
		{"$set", bson.D{{"username", username}}},
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

}
