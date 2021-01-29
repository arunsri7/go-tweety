# go-tweety
  *  #### go-tweety-cli is a cli twitter application. This is a standalone app and it doesn't use the Apis. Check out the [documentation](https://github.com/arunsri7/go-tweety/tree/master/go-tweety-cli) here.
  
  *  #### go-tweety-api has the APIs for similar tweet operations. Here is the [documentation](https://github.com/arunsri7/go-tweety/tree/master/go-tweety-apis)


## Directory tree

     go-tweety
     ├── README.md
     ├── go-tweety-apis
     │   ├── README.md
     │   ├── cmd
     │   │   ├── deleteTweets.go
     │   │   ├── newTweet.go
     │   │   ├── root.go
     │   │   ├── showTweets.go
     │   │   └── signUp.go
     │   ├── go.mod
     │   ├── go.sum
     │   ├── handler
     │   │   ├── admin.go
     │   │   └── users.go
     │   ├── helper
     │   │   └── helper.go
     │   ├── main.go
     │   └── model
     │       └── structs.go
     └── go-tweety-cli
         ├── README.md
         ├── cmd
         │   ├── deleteTweets.go
         │   ├── newTweet.go
         │   ├── root.go
         │   ├── showTweets.go
         │   └── signUp.go
         ├── go.mod
         ├── go.sum
         ├── helper
         │   └── helper.go
         ├── main.go
         └── model
             └── structs.go
