package util

import (
	"gopkg.in/gomail.v2"
)

// 邮箱服务初始化
func init() {
	// 可以在这里初始化一些全局的配置
}

func SendTo(body string, e string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@example.com")
	m.SetHeader("To", e)
	m.SetHeader("Subject", "Subject of the email")
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.example.com", 587, "your-email@example.com", "your-email-password")

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
	}
}
