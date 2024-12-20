package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"time"
)

//rpc 远程过程调用,如何做到想本地调用一样

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	//传输协议:http
	req := HttpRequest.NewRequest()
	// 设置请求超时时间为 10 秒
	req.SetTimeout(10 * time.Second)
	//res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8080/%s?a=%d,b=%d", "add", a, b))
	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8080/%s?a=%d&b=%d", "add", a, b))
	body, _ := res.Body()
	//fmt.Println(string(body))
	rspData := &ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	fmt.Println()
	return rspData.Data
}

func main() {
	fmt.Println(Add(2, 6))

}
