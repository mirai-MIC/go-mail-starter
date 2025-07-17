package mailer

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

var dialer *gomail.Dialer
var globalConfig *MailConfig

// Init 初始化邮件系统（在主程序中调用一次）
func Init(cfg *MailConfig) {
	globalConfig = cfg
	dialer = gomail.NewDialer(cfg.Host, cfg.Port, cfg.User, cfg.Password)
	fmt.Println("Mailer initialized.")
}
