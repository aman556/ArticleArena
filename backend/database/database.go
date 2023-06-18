package database

import (
	//"database/sql"

	//_ "github.com/go-sql-driver/mysql"

	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	//. "github.com/aman556/ArticleArena/backend/database/DAO"
	. "github.com/aman556/ArticleArena/backend/utils"

	// Official 'mongo-go-driver' packages
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var globaldb *sql.DB
var globalClient *mongo.Client

// func InitDB() {
// 	config := NewConfig()

// 	dbServerURL := config.DbUser + ":" + config.DbPass + "@tcp(" + config.DbServiceName + ":" + config.DbHostPort + ")/" + config.DbName
// 	db, err := sql.Open("mysql", dbServerURL)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	globaldb = db
// }

func InitDB() {
	// Declare host and port options to pass to the Connect() method
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	fmt.Println("clientOptions TYPE:", reflect.TypeOf(clientOptions), "\n")

	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	globalClient = client
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func AddUserHandleInDB(userData UserHandles) {
	// query := "INSERT INTO `UserHandles` (`ArticleArenaHandle`, `GeeksforgeeksHandle`, `MediumHandle`, `TutorialpointHandle`) VALUES(?,?,?,?,?)"
	// insert, err := globaldb.Query(query, userData.ArticleArenaHandle, userData.UserHandleList[0].WebsiteHandle, userData.UserHandleList[1].WebsiteHandle, userData.UserHandleList[2].WebsiteHandle)

	// // if there is an error inserting, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // be careful deferring Queries if you are using transactions
	// defer insert.Close()
	// Declare Context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// Access a MongoDB collection through a database
	col := globalClient.Database("article_arena_database").Collection("UserHandles")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	fmt.Println("oneDoc TYPE:", reflect.TypeOf(userData), "\n")

	// InsertOne() method Returns mongo.InsertOneResult
	result, insertErr := col.InsertOne(ctx, userData)
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
		os.Exit(1) // safely exit script on error
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() API result:", result)

		// get the inserted ID string
		newID := result.InsertedID
		fmt.Println("InsertOne() newID:", newID)
		fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
	}
}

func AddUserInfoInDB(userInfo UserInfo) {
	query := "INSERT INTO `UserInfo` (`UserName`, `ArticleArenaHandle`, `Email`, `GithubUrl`, `LinkedinUrl`) VALUES(?,?,?,?,?)"
	insert, err := globaldb.Query(query, userInfo.Name, userInfo.ArticleArenaHandle, userInfo.Email, userInfo.GithubUrl, userInfo.LinkedinUrl)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func SelectUserInfoInDB(userHandle string) {

}

func SelectUserHandleInDB(userHandle string) {

}
