package coreemail

import (
	"fmt"
	"mime"
	"path/filepath"

	"gopkg.in/gomail.v2"
)

type Mail struct {
	SenderAddr   string   // 发件人地址
	SenderName   string   // 发件人名称
	ReceiverAddr []string // 收件人地址，可以有多个收件人
	Subject      string   // 邮件主题
	Text         string   // 正文
	FilePath     string   //附件地址
	Host         string   // 邮件服务器地址
	Port         int      // 邮件服务器端口号
	Username     string   // 用户名
	Password     string   // 密码或授权码
}

// SendMail(&Mail{
// 	SenderAddr:   "sender@163.com",
// 	SenderName:   "senderName",
// 	ReceiverAddr: []string{"receiver@163.com", "receiver@163.com"},
// 	Subject:      "subject",
// 	Text:         "test",
// 	FilePath:	  "C:\title.txt"
// 	Host:         "smtp.163.com",
// 	Port:         25,
// 	Username:     "username@163.com",
// 	Password:     "password",
// })
func SendMail(s *Mail) error {
	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(s.SenderAddr, s.SenderName)}, // 发件人邮箱，发件人名称
		"To":      s.ReceiverAddr,                                // 多个收件人
		"Subject": {s.Subject},                                   // 邮件主题
	})
	m.SetBody("text/plain", s.Text)
	if len(s.FilePath) > 0 {
		fileName := filepath.Base(s.FilePath)
		m.Attach(s.FilePath,
			gomail.Rename(fileName),
			gomail.SetHeader(map[string][]string{
				"Content-Disposition": {
					fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", fileName)),
				},
			}),
		)
	}
	d := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password) // 发送邮件服务器、端口号、发件人账号、发件人密码
	return d.DialAndSend(m)
}
