package sender

import "gopkg.in/gomail.v2"

type EmailSender struct {
	Host          string
	Port          int
	LoginEmail    string
	LoginPassword string
	SenderEmail   string
}

type EmailCommand struct {
	emailSender *EmailSender
}

func NewEmailSender(emailSender *EmailSender) *EmailCommand {
	return &EmailCommand{
		emailSender: emailSender,
	}
}

func (s *EmailCommand) SendEmail(mailAddress []string, senderName, subject, body string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", senderName+"<"+s.emailSender.SenderEmail+">")
	m.SetHeader("To", mailAddress...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(s.emailSender.Host, s.emailSender.Port, s.emailSender.LoginEmail, s.emailSender.LoginPassword)
	err = d.DialAndSend(m)
	return
}
