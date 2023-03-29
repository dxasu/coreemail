err := SendMail(&Mail{
	SenderAddr:   "sender@163.com",
	SenderName:   "senderName",
	ReceiverAddr: []string{"receiver@163.com", "receiver@163.com"},
	Subject:      "subject",
	Text:         "test",
	FilePath:	  "C:\title.txt"
	Host:         "smtp.163.com",
	Port:         25,
	Username:     "username@163.com",
	Password:     "password",
})

if err != nil {
		fmt.Println(err)
	}