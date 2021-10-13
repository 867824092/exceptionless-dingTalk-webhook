package handle

import "github.com/blinkbean/dingtalk"

func SendText(content string) error {
	// 群组机器人的Token
	// 具体用法详见地址：https://github.com/blinkbean/dingtalk
	var dingToken = []string{"6d21f7e26ee2e3a1d8076ce4208eb1b11ad6b1f74d5f243a1267a47aea6aa58e"}
	cli := dingtalk.InitDingTalk(dingToken, ".")
	return cli.SendTextMessage(content)
}
