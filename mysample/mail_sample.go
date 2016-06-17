package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	sses "github.com/tj/go-ses"
)

func main() {
	svc := ses.New(session.New(aws.NewConfig().WithRegion("us-west-2")))
	mailer := sses.SES{svc}

	email := sses.Email{
		From:    "info@hoge.com",
		To:      []string{"tissi0708@gmail.com"},
		Subject: "Hey you!",
		Text:    "Hi",
		HTML:    "<strong>Hi</strong>",
	}

	err := mailer.SendEmail(email)

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
