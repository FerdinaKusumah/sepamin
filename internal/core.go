package internal

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"net/smtp"
)

func Run(opt *Option) {
	var ch = make(chan bool)
	var result = make([]interface{}, opt.Count)
	// auth user
	opt.Auth = smtp.PlainAuth("", opt.User, opt.Password, opt.SmtpHost)

	for i := 0; i < opt.Count; i++ {
		message := []byte(fmt.Sprintf("To: %s\r\n", randomdata.SillyName()) +
			fmt.Sprintf("Subject: %s\r\n", randomdata.SillyName()) + "\r\n" + fmt.Sprintf("%s.\r\n", randomdata.SillyName()),
		)
		go opt.sepaminEmailnyaDong(opt.User, message, i, &result[i], ch)
	}
	// receive chan
	for i := 0; i < opt.Count; i++ {
		<-ch
	}
	// print result
	opt.print(result)
}

func (opt *Option) print(result []interface{}) {
	if opt.Verbose == false {
		fmt.Println("Successfully sent all email")
		return
	}

	for i := 0; i < opt.Count; i++ {
		if v, ok := result[i].(string); ok {
			fmt.Println(v)
		}
	}
}

func (opt *Option) sepaminEmailnyaDong(from string, message []byte, idx int, result *interface{}, ch chan<- bool) {
	var err error
	var status = "successfully sent"
	if err = smtp.SendMail(fmt.Sprintf("%s:%s", opt.SmtpHost, opt.SmtpPort), opt.Auth, from, opt.ListTarget, message); err != nil {
		status = fmt.Sprintf("failed to sent %s", err)
	}
	*result = fmt.Sprintf(`Email %d %s`, idx, status)
	ch <- true
}
