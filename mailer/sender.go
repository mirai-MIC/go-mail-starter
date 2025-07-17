package mailer

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strings"
	"time"
)

// SendHTMLMailFromFile 从文件读取 HTML 模板并发送邮件
func SendHTMLMailFromFile(to, subject, templatePath string, params map[string]string) error {
	if dialer == nil || globalConfig == nil {
		return fmt.Errorf("mailer 未初始化，请先调用 mailer.Init()")
	}

	// 读取模板
	content, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("读取模板失败: %w", err)
	}
	html := string(content)

	// 占位符替换
	for k, v := range params {
		html = strings.ReplaceAll(html, fmt.Sprintf("${%s}", k), v)
	}

	// 构造邮件
	msg := gomail.NewMessage()
	msg.SetHeader("From", globalConfig.User)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", fmt.Sprintf("%s - %s", subject, time.Now().Format("2006-01-02 15:04:05")))
	msg.SetBody("text/html", html)

	// 发送
	return dialer.DialAndSend(msg)
}
