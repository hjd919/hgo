package gom

import (
	"log"

	"gopkg.in/gomail.v2"
)

func NewEmail(conf *EmailConf) *Email {
	// dialer
	dialer := gomail.NewDialer(conf.Host, conf.Port, conf.Username, conf.Password)
	// message
	m := gomail.NewMessage()
	m.SetAddressHeader("From", conf.FromMail, conf.FromName)
	bodyType := "text/html"
	if conf.BodyType != "" {
		bodyType = conf.BodyType
	}

	return &Email{
		dialer:   dialer,
		m:        m,
		conf:     conf,
		bodyType: bodyType,
	}
}

type EmailConf struct {
	FromMail string // 邮件发送方 297538600@qq.com
	FromName string // 邮件发送方昵称 小抿一口
	Host     string // smtp邮件host smtp.qq.com
	Port     int    // smtp邮件端口 465
	Username string // smtp的邮件用户名 hjd
	Password string // smtp的邮件密码 xxx
	BodyType string // opt 邮件内容的类型 "text/html"
}

type Email struct {
	conf     *EmailConf // gom邮件Email配置
	bodyType string     // 邮件内容的类型 "text/html"
	dialer   *gomail.Dialer
	m        *gomail.Message // 邮件Message对象
}
type ToEmail struct {
	Address string // 邮件接收方
	Name    string // 邮件接收方昵称
}

// 给一个用户发送一封邮件
type EmailSendOneParam struct {
	ToEmail ToEmail
	Subject string // 邮件主题
	Body    string // 邮件内容
}

func (t *Email) SendOne(p *EmailSendOneParam) (err error) {
	m := t.m
	d := t.dialer
	m.SetAddressHeader("To", p.ToEmail.Address, p.ToEmail.Name)
	m.SetHeader("Subject", p.Subject)
	m.SetBody(t.bodyType, p.Body)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan") // 抄送
	// m.Attach("/home/Alex/lolcat.jpg") // 附件
	err = d.DialAndSend(m)
	return
}

// 给多个用户发送一封邮件
type EmailSendParam struct {
	ToEmails []ToEmail // 邮件接收方
	Subject  string    // 邮件主题
	Body     string    // 邮件内容
}

func (t *Email) Send(p *EmailSendParam) (err error) {
	m := t.m
	d := t.dialer
	conf := t.conf
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}
	list := p.ToEmails
	for _, r := range list {
		m.SetAddressHeader("From", conf.FromMail, conf.FromName)
		m.SetAddressHeader("To", r.Address, r.Name)
		m.SetHeader("Subject", p.Subject)
		m.SetBody(t.bodyType, p.Body)

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", r.Address, err)
		}
		m.Reset()
	}
	return
}
