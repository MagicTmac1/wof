package tool

type Imp struct {
	Id     string `json:"id"`
	Banner struct {
		W      int `json:"w"`
		H      int `json:"h"`
		Format []struct {
			W int `json:"w"`
			H int `json:"h"`
		} `json:"format"`
		Battr []int `json:"battr"`
		Pos   int   `json:"pos"`
		Api   []int `json:"api"`
		Ext   struct {
			Placementtype string `json:"placementtype"`
		} `json:"ext"`
	} `json:"banner"`
	Displaymanager    string  `json:"displaymanager"`
	Displaymanagerver string  `json:"displaymanagerver"`
	Bidfloor          float64 `json:"bidfloor"`
	Bidfloorcur       string  `json:"bidfloorcur"`
}
type Req struct {
	Id  string `json:"id"`
	Imp []Imp  `bson:"imp"`
	App struct {
		Id        string   `json:"id"`
		Name      string   `json:"name"`
		Cat       []string `json:"cat"`
		Publisher struct {
			Id string `json:"id"`
		} `json:"publisher"`
		Bundle   string `json:"bundle"`
		Storeurl string `json:"storeurl"`
		Ver      string `json:"ver"`
	} `json:"app"`
	Device struct {
		Ua  string `json:"ua"`
		Geo struct {
			Lat      float64 `json:"lat"`
			Lon      float64 `json:"lon"`
			Type     int     `json:"type"`
			Accuracy int     `json:"accuracy"`
			Country  string  `json:"country"`
			Region   string  `json:"region"`
			City     string  `json:"city"`
			Zip      string  `json:"zip"`
		} `json:"geo"`
		Ip             string `json:"ip"`
		Devicetype     int    `json:"devicetype"`
		Make           string `json:"make"`
		Model          string `json:"model"`
		Os             string `json:"os"`
		Osv            string `json:"osv"`
		Language       string `json:"language"`
		Carrier        string `json:"carrier"`
		Connectiontype int    `json:"connectiontype"`
		Ifa            string `json:"ifa"`
	} `json:"device"`
	User struct {
		Id string `json:"id"`
	} `json:"user"`
	At   int      `json:"at"`
	Tmax int      `json:"tmax"`
	Bcat []string `json:"bcat"`
	Regs struct {
		Ext struct {
			Type           string      `json:"Type"`
			CustomTracking interface{} `json:"CustomTracking"`
			Data           interface{} `json:"Data"`
		} `json:"ext"`
	} `json:"regs"`
	Ext struct {
		Exchange string `json:"exchange"`
	} `json:"ext"`
}

type Res struct {
	Id       string    `json:"id"`
	Seatbids []Seatbid `json:"seatbid"`
	Cur      string    `json:"cur"`
}
type Bid struct {
	Id      string   `json:"id"`
	Impid   string   `json:"impid"`
	Price   float64  `json:"price"`
	Adid    string   `json:"adid"`
	Burl    string   `json:"burl"`
	Adm     string   `json:"adm"`
	Adomain []string `json:"adomain"`
	Bundle  string   `json:"bundle"`
	Cid     string   `json:"cid"`
	Crid    string   `json:"crid"`
	Cat     []string `json:"cat"`
	H       int      `json:"h"`
	W       int      `json:"w"`
}
type Seatbid struct {
	Bids []Bid  `json:"bid"`
	Seat string `json:"seat"`
}

//shabi.xyz
type Domain struct {
	Id   string `bson:"_id"`
	Name string `json:"name"`
	Num  int    `json:"num"`
	Ads  []Ad   `bson:"ad"`
}
type Ad struct {
	Adm      string   `bson:"adm"`
	Bundle   string   `bson:"bundle"`
	Platform []string `bson:"platform"`
	Cat      []string `bson:"cat"`
	Burl     string   `bson:"burl"`
}
