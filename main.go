package main

import (
	"email-svc/config"
	"email-svc/sender"
	"fmt"
)

func main() {
	emailSender := sender.NewEmailSender(&sender.EmailSender{
		Host:          config.EmailConfig.Host,
		Port:          config.EmailConfig.Port,
		LoginEmail:    config.EmailConfig.LoginEmail,
		LoginPassword: config.EmailConfig.LoginPassword,
		SenderEmail:   config.EmailConfig.SenderEmail,
	})

	err := emailSender.SendEmail([]string{"mike.liang45@gmail.com"}, "Name", "Subject", "Html Body")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Success")
}
