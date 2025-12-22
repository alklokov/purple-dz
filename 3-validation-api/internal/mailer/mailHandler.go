package mailer

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/request"
	"3-validation-api/pkg/result"
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
	return func(w http.ResponseWriter, r *http.Request) {
		mail, err := request.Decode[MailRequest](r.Body)
		if err == nil {
			err = request.IsValid(mail)
		}
		if err != nil {
			result.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		conf := makeConfirmation(mail.Email)
		err = saveConfirmation(*conf)
		if err != nil {
			result.Json(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		err = m.SendEmail(m.makeConfirmEmail(*conf))
		if err != nil {
			result.Json(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (m *MailHandler) verifyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		fmt.Println("HASH = ", hash)
		if verifyHash(hash) {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}
}

func (m *MailHandler) makeConfirmEmail(conf MailConfirmation) *email.Email {
	url := fmt.Sprintf("<a href=\"http://localhost:8081/verify/%s\">Confirm</a>", conf.Hash)
	return &email.Email{
		From:    m.Sender,
		To:      []string{conf.Email},
		Subject: "Confirm email",
		HTML:    []byte(url),
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
