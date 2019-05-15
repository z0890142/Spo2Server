package DB

import (
	"context"
	"fmt"
	"log"
	"time"
	t "time"

	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	_mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var h t.Duration

func init() {

	h, _ = time.ParseDuration("1h")
}
func Connect_Mongo(addr string) *_mongo.Collection {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin1:821224@localhost:27017/Spo2"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("Mongo")))

	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("Spo2")

	userColl := db.Collection("data1")
	//Insert(userColl, "", 1)
	//UpdateOne
	//Update(userColl, "", 123)
	fmt.Println("Connected to MongoDB!")
	return userColl
}

func Update(collection *_mongo.Collection, deviceid string, spo2 int, bpm int) {
	now := time.Now()
	now.Add(8 * h)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"deviceid": deviceid}
	update_data := bson.M{"spo2": spo2, "bpm": bpm, "time": now.Unix() * 1000}
	data := bson.M{
		"$push": bson.M{"data": update_data},
		"$min":  bson.M{"first": now.Unix() * 1000},
		"$max":  bson.M{"last": now.Unix() * 1000},
		"$inc":  bson.M{"count": 1}}
	if result, err := collection.UpdateOne(ctx, filter, data); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}
}

func Insert(collection *_mongo.Collection, deviceid string, spo2 int, bpm int) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	now := time.Now()
	now.Add(8 * h)
	insert_data := bson.M{
		"deviceid":   deviceid,
		"create_day": now.Unix() * 1000,
		"count":      1,
		"first":      now.Unix() * 1000,
		"last":       now.Unix() * 1000,
		"data":       []bson.M{bson.M{"spo2": spo2, "bpm": bpm, "time": now.Unix() * 1000}}}
	// Insert one
	if result, err := collection.InsertOne(ctx, insert_data); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}
}

func FindOne(collection *_mongo.Collection, deviceid string) bool {
	var result bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	find_data := bson.M{"deviceid": deviceid}
	// Insert onesasd
	collection.FindOne(ctx, find_data).Decode(&result)
	if len(result) != 0 {
		return true
	} else {
		return false

	}
}

func FindOnlyId(collection *_mongo.Collection) interface{} {
	var resultArray []interface{} = make([]interface{}, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	find_data := bson.M{}
	// Insert onesasd
	cur, err := collection.Find(ctx, find_data)
	defer cur.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		resultArray = append(resultArray, result["deviceid"])
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return resultArray

}

func FindOneByApi(collection *_mongo.Collection, deviceid string) interface{} {
	var result bson.M
	// var spo2 []int64 = make([]int64, 0)
	// var bpm []int64 = make([]int64, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	find_data := bson.M{"deviceid": deviceid}
	// Insert onesasd
	collection.FindOne(ctx, find_data).Decode(&result)
	data := result["data"]

	return data

}

func MsgHandler(collection *_mongo.Collection, deviceid string, spo2 int, bpm int) {
	checkObject := FindOne(collection, deviceid)
	if checkObject == true {
		Update(collection, deviceid, spo2, bpm)
	} else {
		Insert(collection, deviceid, spo2, bpm)
	}
}
