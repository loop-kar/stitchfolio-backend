package util

import (
	"os"

	"github.com/imkarthi24/sf-backend/internal/config"

	"github.com/jordan-wright/email"
)

type EmailContent struct {
	To                   []string
	Subject              string
	Message              *string
	HtmlTemplateFileName *string
	TextTemplateFileName *string
	TemplateValueMap     map[string]string
}

type emailType string

const (
	NIL  emailType = "_"
	HTML emailType = "HTML"
	TEXT emailType = "TEXT"
)

func SendEmail(config *config.SMTPConfig, mailContent EmailContent) error {

	// hostAddress := fmt.Sprintf("%s:%d", config.Host, config.Port)

	e := email.NewEmail()
	e.Subject = mailContent.Subject
	e.From = config.UserName
	e.To = mailContent.To

	if config.Override {
		e.To = []string{config.OverrideTo}
	}

	mailtype, text, err := BuildEmailBody(mailContent)
	if err != nil {
		return err
	}

	switch mailtype {
	case NIL:
		e.HTML = []byte(*mailContent.Message)
	case TEXT:
		e.HTML = text
	case HTML:
		e.HTML = text
	}

	// err = e.Send(hostAddress, smtp.PlainAuth("", config.UserName, config.Password, config.Host))
	// if err != nil {
	// 	return err
	// }
	return nil
}

func BuildEmailBody(mailContent EmailContent) (emailType, []byte, error) {

	if !IsNilOrEmptyString(mailContent.HtmlTemplateFileName) {
		htmlTemplateDirectory := "./templates/html_templates/"
		htmlFile := htmlTemplateDirectory + *mailContent.HtmlTemplateFileName
		htmlContent, err := readContentFromFile(htmlFile)
		if err != nil {
			return NIL, nil, err
		}

		htmlContent = ReplaceTemplateValues(htmlContent, mailContent.TemplateValueMap)
		return HTML, htmlContent, nil
	}

	if !IsNilOrEmptyString(mailContent.TextTemplateFileName) {
		htmlTemplateDirectory := "./templates/message_templates/"
		textFile := htmlTemplateDirectory + *mailContent.TextTemplateFileName
		textContent, err := readContentFromFile(textFile)
		if err != nil {
			return NIL, nil, err
		}

		text := ReplaceTemplateValues(textContent, mailContent.TemplateValueMap)
		return TEXT, text, nil
	}

	return NIL, nil, nil
}

func readContentFromFile(fileName string) ([]byte, error) {

	bs, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return bs, nil
}
