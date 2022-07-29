package configs

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// checkErr(err error){

// 	if err != nil{
// 		log.Fatal(err)
// 	}
// }


func ConnectDB() *mongo.Client {
	client , errConn := mongo.NewClient(options.Client().ApplyURI(EnvMongoURL()))

	if errConn != nil{

		log.Fatal("can't connect to mongodb")
	}

	

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	errCtx := client.Connect(ctx)

		if errCtx != nil {
			log.Fatal(errCtx)
		}

	errPing := client.Ping(ctx,nil)
		if errPing != nil {
			log.Fatal(errPing)
		}

		
	
	
	fmt.Println("Connected to MongoDB")


	return client
	


	

}

var DB *mongo.Client = ConnectDB()

	//getting database collections

	func GetCollection(client  *mongo.Client, collectionName string) *mongo.Collection {

		collection := client.Database("soundgadgetAPI").Collection(collectionName)
		
		return collection
	}





