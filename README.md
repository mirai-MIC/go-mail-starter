# 📧 go-mail-starter

一个可复用的 Go 邮件发送模块，支持从本地 HTML 模板文件读取邮件内容，并动态替换 `${xxx}` 占位符。适用于验证码、通知邮件等通用场景。

## ✨ 功能特点

- 支持 SMTP 邮件发送
- 支持从本地 HTML 模板文件读取邮件正文
- 支持多个占位符参数（如 `${code}`、`${username}`）
- 封装初始化方法，支持多项目复用
- 可选导入方式，类比 Java 的 Spring Boot Starter

---

## 📦 安装方式

将本模块作为内部依赖引入，或发布至私有仓库：

```bash
go get github.com/your-org/go-mail-starter
```

或使用 `replace` 引入本地包：

```go
replace github.com/mirai-MIC/go-mail-starter => ../go-mail-starter
```

---

## 🛠️ 使用方法

### 1. 初始化邮件服务

```go
import "your-module-path/mailer"

mailer.Init(&mailer.MailConfig{
    Host: "smtp.example.com",
    Port: 587,
    User: "noreply@example.com",
    Password: "your_smtp_password",
    From: "noreply@example.com", // 可选
})
```

### 2. 发送邮件（从本地 HTML 文件读取模板）

```go
params := map[string]string{
    "username": "主人",
    "code":     "123456",
    "operation": "登录验证",
    "expire":   "10",
    "time":     time.Now().Format("2006-01-02 15:04:05"),
}

err := mailer.SendHTMLMailFromFile(
    "user@example.com",
    "您的操作验证码",
    "template/email.html",
    params,
)

if err != nil {
    log.Println("发送失败:", err)
}
```

---

## 📄 邮件模板示例（`email.html`）

```html
<h2>您好，${username}：</h2>
<p>您正在进行 <strong>${operation}</strong> 操作。</p>
<p>验证码为：</p>
<h1>${code}</h1>
<p>验证码将在 <strong>${expire}</strong> 分钟后失效。</p>
<p>操作时间：${time}</p>
```

支持任意自定义占位符，只需在 HTML 中写 `${key}`，并在参数中传入 `map[key] = value`。

---

## 🔒 安全建议

- 不要将敏感密码（如 SMTP 密码）硬编码在代码中
- 建议通过 `.env` / 配置文件 / 环境变量读取敏感信息

---

## 📁 项目结构建议

```
go-mail-starter/
├── mailer/
│   ├── config.go      # 配置结构体
│   ├── init.go        # 初始化方法
│   ├── sender.go      # 邮件发送逻辑
├── go.mod
├── README.md
```

---

## 📜 License

MIT
