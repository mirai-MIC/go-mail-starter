package main

import (
	"fmt"
	"go-mail-starter/mailer"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	mailer.Init(&mailer.MailConfig{
		Host:     "smtp.example.com",
		Port:     587,
		User:     "you@example.com",
		Password: "your_password",
	})

	params := map[string]string{
		"username":  "主人",
		"code":      "123456",
		"operation": "登录验证",
		"expire":    "10",
		"time":      time.Now().Format("2006-01-02 15:04:05"),
	}

	err := mailer.SendHTMLMailFromFile("3092179918@qq.com", "验证码", "template/email.html", params)
	if err != nil {
		fmt.Println("发送失败：", err)
	}
}
