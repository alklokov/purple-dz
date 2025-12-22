package mailer

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const userFile string = "user.txt"

func makeConfirmation(mail string) *MailConfirmation {
	return &MailConfirmation{
		Email: mail,
		Hash:  fmt.Sprintf("%x", sha256.Sum256([]byte(mail))),
	}
}

func saveConfirmation(conf MailConfirmation) error {
	file, err := os.Create(userFile)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	file.Write(bytes)
	return nil
}

func verifyHash(hash string) bool {
	file, err := os.Open(userFile)
	if err != nil {
		return false
	}
	defer os.Remove(userFile)
	defer file.Close()
	var conf MailConfirmation
	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		return false
	}
	return conf.Hash == strings.Trim(hash, "{}")
}
