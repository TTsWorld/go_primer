package go_mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

var mgoCli *mongo.Client

//var url string = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.6.1"

func initEngine() {
	var err error
	clientOptions := options.Client().ApplyURI(url)

	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	return mgoCli
}

var (
	client     = GetMgoCli()
	db         *mongo.Database
	collection *mongo.Collection
)

func Test0101(t *testing.T) {

	//2.选择数据库 my_db
	db = client.Database("my_db")

	//选择表 my_collection
	collection = db.Collection("my_collection")
	collection = collection
	fmt.Printf("%+v", collection)
}

// 插入某一条数据
func Test0102(t *testing.T) {
	var lr *LogRecord
	iResult, err := client.Database("my_db").Collection("my_collection").InsertOne(context.TODO(), lr)
	if err != nil {
		fmt.Print(err)
		return
	}
	//_id:默认生成一个全局唯一ID
	id := iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID", id.Hex())
}
