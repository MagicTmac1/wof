package tool

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Count(collection *mongo.Collection, num int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("点击次数加一！", num)
		num += 1
		//filter := bson.M{"num": bson.D{{"$type", "int"}}}
		result, err := collection.UpdateOne(context.TODO(), bson.M{"num": bson.D{{"$type", "int"}}}, bson.D{{"$set", bson.M{"num": num}}})
		if err != nil {
			fmt.Println(result)
		}
	}

}
