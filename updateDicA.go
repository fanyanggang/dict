package main

import (
    "fmt"
	"regexp"
	"strconv"
	"io/ioutil"
	"net/http"
	"strings"

	"encoding/json"
)

type DictResp struct {
	Data DictData `json:"data"`
}
type DictData struct {
	Name        string    `json:"f_name"`
	Binary      string    `json:"f_binary"`
	Version     string    `json:"f_version"`
	BidLevel1   string    `json:"f_bid_level1"`
	BidLevel2   string    `json:"f_bid_level2"`
	BidLevel3   string    `json:"f_bid_level3"`
	BidLevel4   string    `json:"f_bid_level4"`
	Creator     string    `json:"f_creator"`
	Responsible string    `json:"f_responsible"`
	Note        string    `json:"f_note"`
	CTime       string `json:"f_ctime"`
	MTime       string `json:"f_mtime"`
	Delete      string    `json:"f_delete"`
	Value       string    `json:"value"`
	Busipath    string    `json:"busipath"`
}

type DictValue struct {
	URL  string `json:"url"`
	MD5  string `json:"md5"`
	Size string `json:"size"`
}

type UpdateResp struct{
	Code  int `json:"code"`
	Msg  string `json:"msg"`
	Data  interface{} `json:"data"`
}

var strTemp = "http://dl.app.qq.com/inews/3327f20191128v5940/TencentNews_88888_v5940.apk 3327feb87e382f52d9365abf7da08497"


var strURL = "http://dl.app.qq.com/inews/3327f20191128v5940/TencentNews_88888_v5940.apk"
//Token = "6c53992c5dd954e374a11681acd3700c"
//User  = "yvanli"

var DictConf = map[int64]map[string]string{
	88888: {
		"url": "test_apk_downloadurl.appnews.com.qqnews", //http://dl.app.qq.com/inews/e735920191209v0000/TencentNews_122_t0000.apk
		"md5": "test_apk_md5.appnews.com.qqnews",         //b67a604a1f2aabc42bc3eba88a0000b2
	},
}

func main(){

	item := strings.Split(strTemp, " ")
	if len(item) != 2 {
		 fmt.Printf("UpdateDic item length illegal:%vd", strTemp)
	}

	reg := regexp.MustCompile("TencentNews_[\\d]+_")              
//	fmt.Printf("%v\n", reg.FindAllString(strURL, -1))

	index := reg.FindIndex([]byte(strURL))
	//fmt.Println("index|", index[0])
	fmt.Println("start:", index[0], ",end:", index[1], strURL[index[0]+12:index[1]-1])
	version := fmt.Sprintf("%s", strURL[index[0]+12:index[1]-1])

	fmt.Println("version|", version)
	channelID,err := strconv.ParseInt(version,10,64)
	if err == nil {
		fmt.Printf("i64: %v\n",channelID)
	}

	// fmt.Printf("update url---------------%v\n", DictConf[channelID]["url"])
	// fmt.Printf("update md5---------------%v\n", DictConf[channelID]["md5"])
	// update(DictConf[channelID]["url"], item[0], strInfo)
	// update(DictConf[channelID]["md5"], item[1], strInfo)
}

func update(sid, value, info string) {
	url := fmt.Sprintf("http://ons.webdev.com/api/updDict?token=2f0aa92d5863e63eafeccdbb0010d5f6&oper_user=roryfan&sid=%s&value=%s&info=%s&encoding=1", sid, value, info)
	resp, err := http.Get(url)
	if err != nil {
        fmt.Println(err)
        return 
    }
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	fmt.Printf("url ---------------%v\n", url)
	fmt.Printf("body ---------------%v\n", string(body))
	updateResp := &UpdateResp{}
	err = json.Unmarshal(body, updateResp)
	if err != nil{
		fmt.Print(err)
	}

	fmt.Printf("---------------%v\n", updateResp)
	return 
}