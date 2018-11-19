package service

import (
	"encoding/json"
	"fmt"
	"sctek.com/typhoon/th-platform-gateway/common"
	"strconv"
)

func callback(d MSG) {
	common.Log.Infoln("yf_manage_message  consumer")
	fmt.Printf("接收到的信息为%q",string(d.Body))
	//发送短息
	go new(MarshalJson).UnmarshalJson(d.Body)
}

func errCallback(d MSG) {
	common.Log.Infoln("errServerQueue consumer")
	fmt.Println(string(d.Body))
}

func dlxCallback(d MSG) {
	common.Log.Infoln("dlxQueue consumer")
	fmt.Println(string(d.Body))
}

func otherCallback(d MSG) {
	common.Log.Infoln("yf_sms_send consumer ")
	fmt.Printf("mq中读到的数据为：%q\r\n",string(d.Body))
	UnmarshalMQBody(d.Body)
}

func  UnmarshalMQBody(body []byte) error {
	common.Log.Infoln("数据库条目 id 解码")
	result := &struct {
		Id string `json:"id"`
	}{}
	err := json.Unmarshal(body, result)
	if err != nil {
		common.Log.Infoln(err)
		return err
	}
	id,err:=strconv.Atoi(result.Id)
	if err!=nil{
		common.Log.Errorln(err)
		return err
	}
	new(LogicService).AboutIdInfo(id)
	return nil
}

