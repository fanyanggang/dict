package main

import (
"net/http"
"fmt"

"strings"
"io/ioutil"
)
func httppost()  {

	data :=`{"msgtype":"markdown","markdown":{"content": "%s"}}`
	req := fmt.Sprintf(data, "test")

	fmt.Printf("---------%s\n", req)

	request, _ := http.NewRequest("POST", "http://in.qyapi.weixin.qq.com/cgi-bin/webhook/send?key=71dddc6a-5407-4fc6-9929-d4df6e95381d", strings.NewReader(req))
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	request.Header.Set("Cache-Control", "no-cache")

    //post数据并接收http响应
    resp,err :=http.DefaultClient.Do(request)
    if err!=nil{
        fmt.Printf("post data error:%v\n",err)
    }else {
        fmt.Println("post a data successful.")
        respBody,_ :=ioutil.ReadAll(resp.Body)
		fmt.Printf("response data:%v\n",string(respBody))
		defer resp.Body.Close()
    }
}

func main(){
	httppost()
}