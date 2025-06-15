package main

import (
	"bytes"
	"html/template"
	"time"

	"github.com/vanng822/go-premailer/premailer"

	mail "github.com/xhit/go-simple-mail/v2"
)

type Mail struct {
	Domain     string
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption string
	FromAddr   string
	FromName   string
}

type Message struct {
	From       string
	FromName   string
	To         string
	Subject    string
	Attachment []string
	Data       any
	DataMap    map[string]any
}

func (m *Mail) getEncryption(s string) mail.Encryption {
	switch s {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}

func (m *Mail) inlineCSS(formattedMessage string) (string, error) {
	opts := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(formattedMessage, &opts)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}

	return html, nil
}

func (m *Mail) buildHTMLMessage(msg Message) (string, error) {
	templateToRender := "./templates/mail.html.gohtml"

	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tmpl bytes.Buffer
	if err = t.ExecuteTemplate(&tmpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	formattedMessage := tmpl.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (m *Mail) buildPlainTextMessage(msg Message) (string, error) {
	templateToRender := "./templates/mail.plain.gohtml"

	t, err := template.New("email-plain").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tmpl bytes.Buffer
	if err = t.ExecuteTemplate(&tmpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	plainMessage := tmpl.String()
	return plainMessage, nil
}

func (m *Mail) sendSMTPMessage(msg Message) error {
	if msg.From == "" {
		msg.From = m.FromAddr
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	data := map[string]any{
		"message": msg.Data,
	}

	msg.DataMap = data

	formattedMessage, err := m.buildHTMLMessage(msg)
	if err != nil {
		return err
	}

	plainMessage, err := m.buildPlainTextMessage(msg)
	if err != nil {
		return err
	}

	srv := mail.NewSMTPClient()
	srv.Host = m.Host
	srv.Port = m.Port
	srv.Username = m.Username
	srv.Password = m.Password
	srv.Encryption = m.getEncryption(m.Encryption)
	srv.KeepAlive = false
	srv.ConnectTimeout = 10 * time.Second
	srv.SendTimeout = 10 * time.Second

	smtpClient, err := srv.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).
		AddTo(msg.To).
		SetSubject(msg.Subject).
		SetBody(mail.TextPlain, plainMessage).
		AddAlternative(mail.TextHTML, formattedMessage)

	for _, attachment := range msg.Attachment {
		email.AddAttachment(attachment)
	}

	err = email.Send(smtpClient)
	if err != nil {
		return err
	}

	return nil
}
