# go-tweety

Go-tweety-CLI is a twitter app which lets you do tweet-operations from the CLI

## Project Setup
    git clone git@github.com:arunsri7/go-tweety.git
    cd go-tweety/go-tweety-cli/
    create a dot env file and add the folowing variables
        * MONGO_URI(your mongo db atlas URI)
        * DB_NAME
        * DOCUMENT_NAME
    Go install

## Usage
    * Use **go-tweety --help** to check all the commands available
    * Use **go-tweety <command name> -- help** to get the spefic command info and usage

## TODO
    * For Admin
        1. Edit user details
        2. Initiate CRUD of tweets on users' behalf

    * For Super-admin:

        1. Approve actions initiated by admin
        2. View logs (e.g access/action/audit log)
        3. Write custom queries to generate insights

            Examples:

            - Post frequency of user X within a timeframe
            - Number of changes requested by Admin P
