package mailer

type MailConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string // 可选：用于展示名称
}
