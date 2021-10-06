package helper

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var (
	mailConf      *viper.Viper
	yourDomain    string
	privateAPIKey string
	senderNameB   string
	senderEmailB  string
	// Logger :
	Logger *zerolog.Logger
)

//MailInit :
func MailInit(config *viper.Viper) {
	mailConf = config

	yourDomain = mailConf.GetString("notification.email.domain")
	privateAPIKey = mailConf.GetString("notification.email.privkey")
	senderNameB = mailConf.GetString("notification.email.sender.from")
	senderEmailB = mailConf.GetString("notification.email.sender.address")
}

//ValidateEmail :
func ValidateEmail(email string) bool {
	v := mailgun.NewEmailValidator(privateAPIKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	vmail, err := v.ValidateEmail(ctx, email, false)
	cancel()
	if err != nil {
		return false
	}
	return vmail.IsValid
}

//SendMailGun :
func SendMailGun(to *mail.Address, subject string, body, tag string, cc []string, bcc []string) (string, error) {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)
	if mg == nil {
		return "", errors.New("failed API key")
	}

	from := mail.Address{Name: senderNameB, Address: senderEmailB}
	sender := from.String()
	recipient := to.String()

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, "", recipient)
	if len(cc) > 0 {
		for _, sc := range cc {
			message.AddCC(sc)
		}
	}
	if len(bcc) > 0 {
		for _, sb := range cc {
			message.AddBCC(sb)
		}
	}

	message.SetHtml(body)

	err := message.AddTag(tag)
	if err != nil {
		Logger.Error().Msgf("%s", err.Error())
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	// Send the message	with a 60 second timeout
	resp, id, err := mg.Send(ctx, message)
	cancel()

	if err != nil {
		Logger.Error().Msgf("%s", err.Error())
		return "", err
	}

	idClean := strings.TrimSuffix(strings.TrimPrefix(id, "<"), ">")

	fmt.Println("send ID ", idClean)
	Logger.Info().Msgf("ID: %s To: %s Subject: %s Resp: %s\n", id, to.Address, subject, resp)
	return idClean, nil
}
