package email

import (
	"os"

	"github.com/imkarthi24/sf-backend/pkg/util"
	mail "github.com/jordan-wright/email"
)

func (s *emailService) SendEmail(mailContent EmailContent) error {

	// hostAddress := fmt.Sprintf("%s:%d", config.Host, config.Port)

	e := mail.NewEmail()
	e.Subject = mailContent.Subject
	e.From = s.config.UserName
	e.To = mailContent.To

	if s.config.Override {
		e.To = []string{s.config.OverrideTo}
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

	if !util.IsNilOrEmptyString(mailContent.HtmlTemplateFileName) {
		htmlTemplateDirectory := "./templates/html_templates/"
		htmlFile := htmlTemplateDirectory + *mailContent.HtmlTemplateFileName
		htmlContent, err := readContentFromFile(htmlFile)
		if err != nil {
			return NIL, nil, err
		}

		htmlContent = util.ReplaceTemplateValues(htmlContent, mailContent.TemplateValueMap)
		return HTML, htmlContent, nil
	}

	if !util.IsNilOrEmptyString(mailContent.TextTemplateFileName) {
		htmlTemplateDirectory := "./templates/message_templates/"
		textFile := htmlTemplateDirectory + *mailContent.TextTemplateFileName
		textContent, err := readContentFromFile(textFile)
		if err != nil {
			return NIL, nil, err
		}

		text := util.ReplaceTemplateValues(textContent, mailContent.TemplateValueMap)
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
