package internal

import (
	"flag"
	"net/smtp"
	"os"
	"strings"
)

type Option struct {
	User        string
	Password    string
	TargetEmail string
	Count       int
	From        string
	SmtpHost    string
	SmtpPort    string
	Verbose     bool
	ListTarget  []string
	Auth        smtp.Auth
}

func ParseOption() *Option {
	var opt = new(Option)

	flag.StringVar(&opt.User, "user", "", "Username email")
	flag.StringVar(&opt.User, "u", "", "Username email")

	flag.StringVar(&opt.From, "from", "", "Sender user")
	flag.StringVar(&opt.From, "f", "", "Sender user")

	flag.StringVar(&opt.Password, "password", "", "")
	flag.StringVar(&opt.Password, "pwd", "", "")

	flag.StringVar(&opt.TargetEmail, "target", "", "Email target ex: example@gmail.com")
	flag.StringVar(&opt.TargetEmail, "t", "", "Email target ex: example@gmail.com")

	flag.IntVar(&opt.Count, "count", 5, "How much you want to spam email")
	flag.IntVar(&opt.Count, "c", 5, "How much you want to spam email")

	flag.StringVar(&opt.SmtpHost, "host", "smtp.gmail.com", "Smtp host email, default use google host")
	flag.StringVar(&opt.SmtpHost, "h", "smtp.gmail.com", "Smtp host email, default use google host")

	flag.StringVar(&opt.SmtpPort, "port", "587", "Smtp port email, default use google port")
	flag.StringVar(&opt.SmtpPort, "p", "587", "Smtp port email, default use google port")

	flag.BoolVar(&opt.Verbose, "verbose", false, "Print verbose message")
	flag.BoolVar(&opt.Verbose, "v", false, "Print verbose message")

	flag.Parse()

	if opt.User == "" || opt.Password == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if opt.TargetEmail == "" {
		opt.TargetEmail = opt.User
	}

	opt.ListTarget = []string{opt.TargetEmail}
	if strings.Contains(opt.TargetEmail, ",") {
		// remove white space
		opt.TargetEmail = strings.ReplaceAll(opt.TargetEmail, " ", "")
		opt.ListTarget = strings.Split(opt.TargetEmail, ",")
	}

	return opt
}
