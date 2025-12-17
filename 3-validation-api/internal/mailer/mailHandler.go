package mailer

import (
	"3-validation-api/configs"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// Структура для конфигурации SMTP
type MailHandler struct {
	*configs.SMTPConfig
}

func MakeMailHandler(router *http.ServeMux, config *configs.Config) {
	m := &MailHandler{
		SMTPConfig: &config.SMTPConf,
	}
	router.HandleFunc("POST /send", m.sendHandler())
	router.HandleFunc("GET /verify/{hash}", m.verifyHandler())
}

func (m *MailHandler) sendHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		err := m.SendEmail(makeTestEmail())
		if err != nil {
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func (m *MailHandler) verifyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("verified"))
	}
}

func makeTestEmail() *email.Email {
	return &email.Email{
		From:    "Sender <sender@mail.ru>",
		To:      []string{"klokov@ravelinspb.ru"},
		Subject: "Test letter",
		Text:    []byte("This is a test letter"),
	}
}

// Функция отправки письма
func (m *MailHandler) SendEmail(e *email.Email) error {
	// Настраиваем аутентификацию
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)

	// Конфигурация TLS для самоподписанных сертификатов
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // ВНИМАНИЕ: только для тестов!
		ServerName:         m.Host,
	}

	// Адрес сервера
	addr := fmt.Sprintf("%s:%d", m.Host, m.Port)

	// Отправляем с поддержкой STARTTLS
	return e.SendWithStartTLS(addr, auth, tlsConfig)
}
