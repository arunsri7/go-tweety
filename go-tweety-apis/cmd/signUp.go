package cmd

import (
	"context"
	"go-tweety/helper"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// signUpCmd represents the signUp command
var signUpCmd = &cobra.Command{
	Use:   "signUp <username> <password> <consumerKey> <consumerSecret> <accessToken> <accessSecret>",
	Short: "Adds your username,password,api keys and access token keys to the db in encrypted form for signup ",
	Run: func(cmd *cobra.Command, args []string) {
		password := helper.Encrypt([]byte(args[1]), "password")
		consumerKey := helper.Encrypt([]byte(args[2]), "apiKey")
		consumerSecret := helper.Encrypt([]byte(args[3]), "apiSecret")
		accessToken := helper.Encrypt([]byte(args[4]), "accessToken")
		accessSecret := helper.Encrypt([]byte(args[5]), "accessSecret")
		signUp(args[0], password, consumerKey, consumerSecret, accessToken, accessSecret)
	},
}

func init() {
	rootCmd.AddCommand(signUpCmd)
}

func signUp(username string, password []byte, consumerKey []byte, consumerSecret []byte,
	accessToken []byte, accessSecret []byte) {

	//Load Env file
	err := godotenv.Load(".env")

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

	err = client.Connect(ctx)

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
