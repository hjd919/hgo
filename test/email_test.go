package test

import (
	"testing"

	"github.com/hjd919/gom"
)

func TestEmail_Send(t *testing.T) {
	config := &gom.EmailConf{
		FromMail: "297538600@qq.com",
		Host:     "smtp.qq.com",
		Port:     465,
		FromName: "小抿一口",
		Username: "297538600@qq.com",
		Password: "--",
	}
	e := gom.NewEmail(config)
	// toEmails := []ToEmail{}
	// toEmails = append(toEmails, ToEmail{Address: "hjdweapp@163.com", Name: "hjdweapp"})
	p := &gom.EmailSendParam{
		ToEmails: []gom.ToEmail{{Address: "297538600@qq.com", Name: "hjdweapp"}},
		Subject:  "主题subject",
		Body:     "多个邮件人",
	}
	if err := e.Send(p); err != nil {
		t.Errorf("Email.Send() error = %v", err)
	}
}

func TestEmail_SendOne(t *testing.T) {
	config := &gom.EmailConf{
		FromMail: "297538600@qq.com",
		Host:     "smtp.qq.com",
		Port:     465,
		FromName: "小抿一口",
		Username: "297538600@qq.com",
		Password: "--",
	}
	e := gom.NewEmail(config)
	p := &gom.EmailSendOneParam{
		ToEmail: gom.ToEmail{Address: "hjdweapp@163.com", Name: "hjdweapp"},
		Subject: "主题subject",
		Body:    "单个邮件人",
	}
	if err := e.SendOne(p); err != nil {
		t.Errorf("Email.Send() error = %v", err)
	}
}
