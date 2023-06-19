package main

import (
	"github.com/gen2brain/beeep"
	"testing"
)

func TestNotify(t *testing.T) {
	err := beeep.Notify("短信", "【美团买菜】1352（您的手机登录验证码，请完成验证），如非本人操作，请忽略本短信。", "icon.jpeg")
	if err != nil {
		panic(err)
	}
}
