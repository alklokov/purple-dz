package mailer

type MailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type MailConfirmation struct {
	Email string `json:"email" validate:"required,email"`
	Hash  string `json:"hash"`
}
