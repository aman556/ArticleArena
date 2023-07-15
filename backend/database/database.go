package database

import (
	//"database/sql"

	//_ "github.com/go-sql-driver/mysql"

	"context"
	"fmt"
	"reflect"
	"time"

	// . "github.com/aman556/ArticleArena/backend/database/DAO"
	. "github.com/aman556/ArticleArena/backend/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, context.CancelFunc) {
	// config := NewConfig()
	// credential := options.Credential{
	// 	AuthSource: config.DbName,
	// 	Username:   config.DbUser,
	// 	Password:   config.DbPass,
	// }

	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://AmanSharma:8512810555@cluster0.wsuuloc.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, cancel
}

func AddUserHandleInDB(userData UserHandles) {
	// Declare Context type object for managing multiple API requests
	ctx, CancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer CancelFunc()

	// Access a MongoDB collection through a database
	client, _ := ConnectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	col := client.Database("article_arena_database").Collection("UserHandles")
	fmt.Println("Collection type:", reflect.TypeOf(col))
	fmt.Println("oneDoc TYPE:", reflect.TypeOf(userData))

	// InsertOne() method Returns mongo.InsertOneResult
	result, insertErr := col.InsertOne(ctx, userData)
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
		panic(insertErr)
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
	// Declare Context type object for managing multiple API requests
	ctx, CancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer CancelFunc()

	// Access a MongoDB collection through a database
	client, _ := ConnectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	col := client.Database("article_arena_database").Collection("UserInfo")
	fmt.Println("Collection type:", reflect.TypeOf(col))
	fmt.Println("oneDoc TYPE:", reflect.TypeOf(userInfo))

	// InsertOne() method Returns mongo.InsertOneResult
	result, insertErr := col.InsertOne(ctx, userInfo)
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
		panic(insertErr)
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() API result:", result)

		// get the inserted ID string
		newID := result.InsertedID
		fmt.Println("InsertOne() newID:", newID)
		fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
	}
}

func SelectUserInfoInDB(userHandle string) {

}

func SelectUserHandleInDB(userHandle string) {

}
