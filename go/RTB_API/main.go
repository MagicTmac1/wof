package main

import (
	"RTB_API/tool"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
	"time"
)

//接收request中的数据
func AcceptReqData(collection *mongo.Collection) (r tool.Res) {
	var result tool.Req
	err := collection.FindOne(context.TODO(), bson.D{{}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	//将result数组从接收到的数据中提取出来需要的数据
	r.Cur = result.Imp[0].Bidfloorcur
	r.Id = result.Id
	r.Seatbids = append(r.Seatbids, tool.Seatbid{
		Seat: "",
		Bids: []tool.Bid{},
	})
	r.Seatbids[0].Bids = append(r.Seatbids[0].Bids, tool.Bid{})
	r.Seatbids[0].Bids[0].Impid = result.Id

	r.Seatbids[0].Bids[0].Price = result.Imp[0].Bidfloor
	r.Seatbids[0].Bids[0].H = result.Imp[0].Banner.H
	r.Seatbids[0].Bids[0].W = result.Imp[0].Banner.W
	r.Seatbids[0].Bids[0].Cid = time.Now().String() + "wof"
	r.Seatbids[0].Bids[0].Crid = time.Now().String() + result.Device.Os + result.Device.Osv
	r.Seatbids[0].Bids[0].Bundle = result.Device.Os
	return
}
func CompleteRes(res tool.Res, collection *mongo.Collection) (tool.Res, int) {
	var ad tool.Domain
	err := collection.FindOne(context.TODO(), bson.D{{}}).Decode(&ad)
	res.Seatbids[0].Bids[0].Adm = ad.Ads[0].Adm
	res.Seatbids[0].Bids[0].Adomain = append(res.Seatbids[0].Bids[0].Adomain, ad.Name)
	res.Seatbids[0].Bids[0].Cat = ad.Ads[0].Cat
	res.Seatbids[0].Bids[0].Id = ad.Id
	res.Seatbids[0].Bids[0].Adid = ad.Id
	res.Seatbids[0].Bids[0].Burl = "127.0.0.1:9000/count"
	if err != nil {
		fmt.Printf("%v", err)
	}
	for _, p := range ad.Ads[0].Platform {
		if p == res.Seatbids[0].Bids[0].Bundle {
			res.Seatbids[0].Bids[0].Bundle = ad.Ads[0].Bundle + "." + p
		}
	}
	res.Seatbids[0].Seat = collection.Name()
	num := ad.Num
	return res, num
}

//闭包传入数据
func HandleFuncGet(Data tool.Res) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//先获取http筛选条件
		w.Header().Set("content-type", "application/json")
		JsonData, _ := json.Marshal(Data)
		io.WriteString(w, string(JsonData))
	}
}

func main() {
	//连接mongodb
	client := tool.ConnectDB()
	//defer client.Disconnect(context.TODO())
	collectionReq := tool.ConnectCollection(client, "requset")
	//接收部分response数据
	responseDate := AcceptReqData(collectionReq)
	name := "shabi.xyz"
	collectionAd := tool.ConnectCollection(client, name)
	data, num := CompleteRes(responseDate, collectionAd)
	//创建HTTP服务

	http.HandleFunc("/getdata", HandleFuncGet(data))
	http.HandleFunc("/count", tool.Count(collectionAd, num))
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
