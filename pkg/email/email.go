package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

func NewEmail(SMTPInfo *SMTPInfo) *Email {
	return &Email{SMTPInfo: SMTPInfo}
}

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

// SendEmail 指定发送人切片以及主题和内容
func (e *Email) SendEmail(to []string, subject, body string) error {
	message := gomail.NewMessage() //设置发送信息
	message.SetHeader("From", e.From)
	message.SetHeader("To", to...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	//忽略ssl加密
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: e.IsSSL,
	}
	return dialer.DialAndSend(message) //发送内容
}
