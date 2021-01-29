package test

import (
	"testing"

	"github.com/hjd919/gom"
)

func TestSms_Send(t *testing.T) {
	e := gom.NewSms(&gom.SmsConf{
		AccessKeyId:     "-",
		AccessKeySecret: "--",
		SignName:        "师生汇",
		TemplateCode:    "SMS_200185878",
	})
	if err := e.Send(&gom.SmsSendParam{
		Phones: []string{"18500223089"},
		TemplateParam: map[string]interface{}{
			"code": 12345,
		},
	}); err != nil {
		t.Errorf("Sms.Send() error = %v", err)
	}
}
