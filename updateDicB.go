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


//Token = "6c53992c5dd954e374a11681acd3700c"
//User  = "yvanli"

var DictConf = map[int64]map[string]string{
	88888: {
		"url": "test_apk_downloadurl.appnews.com.qqnews", //http://dl.app.qq.com/inews/e735920191209v0000/TencentNews_122_t0000.apk
		"md5": "test_apk_md5.appnews.com.qqnews",         //b67a604a1f2aabc42bc3eba88a0000b2
	},
}

//dict_88888roryfanqqnews   http://dl.app.qq.com/inews/e735920191209v0000/TencentNews_122_t0000.apk
//apk.cfg.88888     {"url":"http:\/\/view.inews.qq.com\/newsDownLoad?refer=biznew&src=88888roryfanqqnews&by=dict","md5":"68981970dab50f409dd43837844d18aa","size":"20M"}


func main(){
	
	// value := DictValue{
	// 	URL:"http://view.inews.qq.com/newsDownLoad?refer=biznew&src=4099newsflowf&by=dict",
	// 	MD5: "666666666666666666666666",
	// 	Size:"20M",
	// }
	// req, err := json.Marshal(value)
	// if err != nil{
	// 	fmt.Println(err)
	// }

    strInfo := fmt.Sprint("渠道包信息测试")
	//get(string(req),strInfo)

	item := strings.Split(strTemp, " ")
	if len(item) != 2 {
		 fmt.Printf("UpdateDic item length illegal:%vd", strTemp)
	}

	reg := regexp.MustCompile("TencentNews_[\\d]+_")              
	fmt.Printf("%q\n", reg.FindAllString(strTemp, -1))

	index := reg.FindIndex([]byte(strTemp))
	//fmt.Println("start:", index[0], ",end:", index[1], strTemp[index[0]+12:index[1]-1])
	version := fmt.Sprintf("%s", strTemp[index[0]+12:index[1]-1])

	channelID,err := strconv.ParseInt(version,10,64)
	if err == nil {
		fmt.Printf("i64: %v\n",channelID)
	}

	dictKey := fmt.Sprintf("apk.cfg.%d", channelID)
	fmt.Printf("dictKey url---------------%v\n", dictKey)
	get(dictKey, item[1], item[0], strInfo)
}



func get(sid, md5, valueUrl, info string) DictValue{

	url := fmt.Sprintf("http://ons.webdev.com/api/getDictInfo?token=2f0aa92d5863e63eafeccdbb0010d5f6&oper_user=roryfan&sid=%s", sid)
	resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return DictValue{}
    }
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	date := new(DictResp)
	err = json.Unmarshal(body, date)

	temp := DictValue{}
	err = json.Unmarshal([]byte(date.Data.Value), &temp)
	temp.MD5 = md5
	temp.Size = "999M"
	req, err := json.Marshal(temp)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("value md5---------:%v", req)
	update(sid, string(req), info)

	dictUpdat, err := convertUrlQuery(temp.URL)
	if err != nil {
		fmt.Println(err)
	}


	update("dict_" + dictUpdat, valueUrl, info)
	fmt.Println(temp)
	return DictValue{}
}

func convertUrlQuery(query string) (string, error) {

	var dict string

	item := strings.Split(query, "&")
	for _, v := range item {
		if strings.Contains(v, "src=") {
			dict = fmt.Sprintf("%s", v[4:])
		}
	}

	return dict, nil
}

func update(sid, value, info string) {
	url := fmt.Sprintf("http://ons.webdev.com/api/updDict?token=2f0aa92d5863e63eafeccdbb0010d5f6&oper_user=roryfan&sid=%s&value=%s&info=%s&encoding=1", sid, value, info)
	fmt.Println(url)
	
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