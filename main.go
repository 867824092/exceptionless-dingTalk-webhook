package main

import (
	"encoding/json"
	"exceptionless.dingTalk/handle"
	model "exceptionless.dingTalk/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	address = "192.168.31.211:7000"
)

// exceptionless 接入WebHook发送钉钉群组机器人通知
func main() {
	http.HandleFunc("/dingTalk/robot", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("read body err:", err.Error())
			return
		}
		go func(param []byte) {
			entity := &model.ExceptionlessWebHookModel{}
			if body != nil {
				if er := json.Unmarshal(param, &entity); er != nil {
					log.Println("解析json错误：", er.Error())
					return
				}
				if entity != nil {
					str, _ := json.Marshal(*entity)
					if str != nil {
						// 文本格式
						handle.SendText(fmt.Sprintf("监控报警:", string(str)))
					}
				}
			}
		}(body)
	})
	http.ListenAndServe(address, nil)
}
