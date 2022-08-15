package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
)

//数据参数
type ApiPara struct {
	Qid        string `bson:"qid"`
	Status     int    `bson:"status"`
	Errmsg     string `bson:"errmsg"`
	Token      string `bson:"token"`
	Timetake   int    `bson:"timetake"`
	Total_page int    `bson:"total_page"`
	Total_num  int    `bson:"total_num"`
	Last_page  bool   `bson:"last_page"`
	Data       []Data `bson:"data"`
}
type Data struct {
	OfferId         string   `bson:"offer_id"`
	OfferName       string   `bson:"offer_name"`
	AppIcon         string   `bson:"app_icon"`
	ConvEvent       string   `bson:"conv_event "`
	Platform        string   `bson:"platform"`
	CountryCode     string   `bson:"country_code"`
	PriceMode       string   `bson:"price_mode"`
	Payout          float64  `bson:"payout"`
	Currency        string   `bson:"currency"`
	DailyConvCap    int      `bson:"daily_conv_cap"`
	DailyClickCap   int      `bson:"daily_click_cap"`
	DailyImpCap     int      `bson:"daily_imp_cap"`
	Package         string   `bson:"package"`
	AppName         string   `bson:"app_name"`
	Category        string   `bson:"category"`
	PreviewUrl      string   `bson:"preview_url"`
	Tags            []string `bson:"tags"`
	Stability       int      `bson:"stability"`
	RecentStability int      `bson:"recent_stability"`
	Star            float64  `bson:"star"`
	Direct          bool     `bson:"direct"`
	IabCats         []string `bson:"iab_cats"`
	Adomain         string   `bson:"adomain"`
	ClickUrl        string   `bson:"click_url"`
	ImpressionUrl   string   `bson:"impression_url"`
	Target          struct {
		OsVersion struct {
			Min string `bson:"min"`
		} `bson:"os_version"`
	} `bson:"target"`
	Performance struct {
		D30 struct {
			ClickNum int     `bson:"click_num"`
			ConvNum  int     `bson:"conv_num"`
			Revenue  float64 `bson:"revenue"`
		} `bson:"d30"`
		D7 struct {
			ClickNum int     `bson:"click_num"`
			ConvNum  int     `bson:"conv_num"`
			Revenue  float64 `bson:"revenue"`
		} `bson:"d7"`
		D3 struct {
			ClickNum int     `bson:"click_num"`
			ConvNum  int     `bson:"conv_num"`
			Revenue  float64 `bson:"revenue"`
		} `bson:"d3"`
		D1 struct {
			ClickNum int     `bson:"click_num"`
			ConvNum  int     `bson:"conv_num"`
			Revenue  float64 `bson:"revenue"`
		} `bson:"d1"`
	} `bson:"performance"`
	FrequencyCtrl  float64 `bson:"frequency_ctrl"`
	AttributionSys string  `bson:"attribution_sys"`
	Hops           int     `bson:"hops"`
	Status         string  `bson:"status"`
	Desc           string  `bson:"desc"`
	Kpi            string  `bson:"kpi"`
	AssignDt       string  `bson:"assign_dt"`
	CreateTs       int     `bson:"create_ts"`
	ModifyTs       int     `bson:"modify_ts"`
}

//连接mongodb
func ConnectDB() (*mongo.Collection, *mongo.Client) {
	//建立连接
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var ctx = context.TODO()
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	//连接到指定db的collection
	collection := client.Database("oyc_task").Collection("api_data")
	return collection, client
}

//接收数据
func AcceptData(collection *mongo.Collection) []Data {
	var result ApiPara
	//定义filter筛选
	//filter := bson.D{{"qid", bson.D{{"$type", "string"}}}}
	err := collection.FindOne(context.TODO(), bson.D{{}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	//将data数组从接收到的数据中提取出来需要的数据
	return result.Data
}

//冒泡排序

func BubbleSort(arr []Data) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i].Payout > arr[j].Payout {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

//闭包传入数据
func HandleFuncGet(AllData []Data) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//先获取http筛选条件
		w.Header().Set("content-type", "application/json")
		values := r.URL.Query()
		//range循环返回出符合条件的offer
		OfferList := []Data{}
		for i := 0; i < len(AllData); i++ {
			if AllData[i].Platform == values.Get("platform") && AllData[i].CountryCode == values.Get("country_code") {
				OfferList = append(OfferList, AllData[i])
			}
		}
		//进行判断，如果满足条件的只有一个广告，那就直接放回这条广告消息，如果有多条广告的话就按照单价进行排序，返回收益最高的那一条广告。
		if len(OfferList) == 0 {
			io.WriteString(w, string(""))
		} else if len(OfferList) == 1 {
			JsonData, _ := json.Marshal(OfferList)
			io.WriteString(w, string(JsonData))
		} else {
			BubbleSort(OfferList)
			JsonData, _ := json.Marshal(OfferList[len(OfferList)-1])
			io.WriteString(w, string(JsonData))
		}
	}
}
func main() {
	//连接mongodb
	
	collection, client := ConnectDB()
	defer client.Disconnect(context.TODO())
	//接收bson数据
	bsonDate := AcceptData(collection)
	//创建HTTP服务
	http.HandleFunc("/getdata", HandleFuncGet(bsonDate))
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Println(err)
	}

}
