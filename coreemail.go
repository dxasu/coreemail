package coreemail

import (
	"fmt"
	"mime"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	SenderAddr   string   `validate:"required,email"`         // 必填且符合邮箱格式
	SenderName   string   `validate:"required,min=1,max=20"`  // 必填，长度1-50字符
	ReceiverAddr []string `validate:"required,dive,email"`    // 必填，每个元素需为邮箱格式
	Subject      string   `validate:"required,min=1,max=100"` // 必填，长度1-100字符
	Text         string   // 正文
	FilePaths    []string `validate:"dive,file"`                // 每个附件路径需为有效文件路径[5,7](@ref)
	Host         string   `validate:"required,hostname|ipv4"`   // 必填，需为域名或IPv4地址[1,7](@ref)
	Port         int      `validate:"required,min=1,max=65535"` // 必填，端口号范围1-65535
	Username     string   `validate:"required"`                 // 必填
	Password     string   `validate:"required"`                 // 必填
}

func (m *Mail) Validate() error {
	validate := validator.New()
	err := validate.Struct(m)
	return err
}

//	SendMail(&Mail{
//		SenderAddr:   "sender@163.com",
//		SenderName:   "senderName",
//		ReceiverAddr: []string{"receiver@163.com", "receiver@163.com"},
//		Subject:      "subject",
//		Text:         "test",
//		FilePaths:	  []string("C:\title.txt")
//		Host:         "smtp.163.com",
//		Port:         25,
//		Username:     "username@163.com",
//		Password:     "password",
//	})
func SendMail(s *Mail) error {
	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(s.SenderAddr, s.SenderName)}, // 发件人邮箱，发件人名称
		"To":      s.ReceiverAddr,                                // 多个收件人
		"Subject": {s.Subject},                                   // 邮件主题
	})
	m.SetBody("text/plain", s.Text)
	if len(s.FilePaths) > 0 {
		for _, filePath := range s.FilePaths {
			fileName := filepath.Base(filePath)
			m.Attach(filePath,
				gomail.Rename(fileName),
				gomail.SetHeader(map[string][]string{
					"Content-Disposition": {
						fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", fileName)),
					},
				}),
			)
		}
	}
	d := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password) // 发送邮件服务器、端口号、发件人账号、发件人密码
	return d.DialAndSend(m)
}
