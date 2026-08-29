package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/core/port"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/core/util"
	"gopkg.in/gomail.v2"
)

// confirmationTmpl is parsed once at init time, not on every email send.
var confirmationTmpl = template.Must(template.New("confirmation").Parse(`
	<h2>Email Confirmation</h2>
	<p>Please confirm your email by clicking the link below:</p>
	<p><a href="{{.URL}}">Confirm Email</a></p>
	<p>This link will expire in 15 minutes.</p>
`))

type NotifService struct {
	repo port.NotifRepository
	conf *config.SMTP
}

func NewNotifService(repo port.NotifRepository, conf *config.SMTP) *NotifService {
	return &NotifService{repo, conf}
}

func (s *NotifService) ReceiveNotif(ctx context.Context) (<-chan amqp.Delivery, error) {
	return s.repo.ReceiveNotif(ctx)
}

func (s *NotifService) SendConfirmationEmail(ctx context.Context, msg []byte) error {
	var data map[string]string
	if err := util.Deserialize(msg, &data); err != nil {
		return fmt.Errorf("notif: failed to deserialize message: %w", err)
	}

	email, ok := data["email"]
	if !ok || email == "" {
		return fmt.Errorf("notif: missing email field in message")
	}

	token, ok := data["confirmation_token"]
	if !ok || token == "" {
		return fmt.Errorf("notif: missing confirmation_token field in message")
	}

	url := util.GenerateConfirmationURL(token)

	// render HTML body via template: avoids XSS from raw string concatenation.
	var body bytes.Buffer
	if err := confirmationTmpl.Execute(&body, map[string]string{"URL": url}); err != nil {
		return fmt.Errorf("notif: failed to render email template: %w", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.conf.SenderEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "User registration confirmation")
	m.SetBody("text/html", body.String())

	port, err := strconv.Atoi(s.conf.Port)
	if err != nil {
		return fmt.Errorf("notif: failed to parse SMTP port: %w", err)
	}
	d := gomail.NewDialer(s.conf.Host, port, s.conf.SenderEmail, s.conf.Password)
	d.SSL = false // STARTTLS on port 587

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("notif: failed to send email to %s: %w", email, err)
	}

	return nil
}
